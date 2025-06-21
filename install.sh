#!/bin/bash

set -e  # Exit on any error

if [ "$EUID" -ne 0 ]; then
    echo "Please run as root"
    exit 1
fi

echo "Installing Home Control Center..."

# Set defaults
DB_PATH="/var/lib/home-control-center/home-control.db"
LISTEN_ADDRESS="0.0.0.0"
LISTEN_PORT=9040

echo ""
echo "Using database: $DB_PATH"
echo "API endpoint: http://$LISTEN_ADDRESS:$LISTEN_PORT"
echo ""

# Get the actual user (not root when using sudo)
ACTUAL_USER="${SUDO_USER:-$USER}"
if [ "$ACTUAL_USER" = "root" ]; then
    echo "Warning: Running as actual root user. User permissions may not work correctly."
fi

# Create installation directory
INSTALL_DIR="/usr/bin/home-control-center"
CONFIG_DIR="/etc/home-control-center"
DATA_DIR=$(dirname "$DB_PATH")
mkdir -p "$INSTALL_DIR"
mkdir -p "$CONFIG_DIR"
mkdir -p "$DATA_DIR"

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        BINARY_NAME="home-control-center-linux-amd64"
        ;;
    aarch64)
        BINARY_NAME="home-control-center-linux-arm64"
        ;;
    armv6l)
        BINARY_NAME="home-control-center-linux-arm6"
        ;;
    armv7l)
        BINARY_NAME="home-control-center-linux-arm7"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        echo "Supported: x86_64, aarch64, armv6l, armv7l"
        exit 1
        ;;
esac

echo "Detected architecture: $ARCH (using $BINARY_NAME)"

# Get latest release info
echo "Fetching latest release..."
LATEST_URL=$(curl -s https://api.github.com/repos/NotCoffee418/home-control-center/releases/latest | grep "browser_download_url.*$BINARY_NAME" | cut -d '"' -f 4)

if [ -z "$LATEST_URL" ]; then
    echo "Error: Could not find download URL for $BINARY_NAME"
    exit 1
fi

echo "Downloading from: $LATEST_URL"

# Stop service if it exists (for updates)
echo "Stopping existing service if running..."
systemctl stop home-control-center.service 2>/dev/null || true
sleep 1

# Download the binary
curl -L -o "$INSTALL_DIR/home-control-center" "$LATEST_URL"
chmod +x "$INSTALL_DIR/home-control-center"

echo "Binary installed to $INSTALL_DIR/home-control-center"

# Create config file if it doesn't exist
if [ ! -f "$CONFIG_DIR/config.toml" ]; then
CONFIG_FILE="$CONFIG_DIR/config.toml"
cat > "$CONFIG_FILE" << EOF
# See README.md for more info on the config file
database_path = "$DB_PATH"
listen_address = "$LISTEN_ADDRESS"
listen_port = $LISTEN_PORT
EOF
fi

echo "Created config file at $CONFIG_FILE"

# Create systemd service
SERVICE_FILE="/etc/systemd/system/home-control-center.service"
cat > "$SERVICE_FILE" << EOF
[Unit]
Description=Home Control Center
After=network.target

[Service]
Type=simple
User=root
ExecStart=$INSTALL_DIR/home-control-center
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF

echo "Created systemd service"

# Reload systemd, enable and start service
echo "Starting service..."
systemctl daemon-reload
systemctl enable home-control-center.service
systemctl restart home-control-center.service

# Wait a bit for service to start
echo "Waiting for service to start..."
sleep 5

# Test the service
echo "Testing service..."
if command -v python3 &> /dev/null; then
    if curl -s http://localhost:9040/api/health | python3 -m json.tool > /dev/null 2>&1; then
        echo "✅ Service is running and responding with valid JSON!"
    else
        echo "❌ Service test failed. Check status with:"
        echo "systemctl status home-control-center"
        echo "journalctl -u home-control-center -f"
        echo "You may need to update the config file at /etc/home-control-center/config.toml"
        exit 1
    fi
else
    echo "⚠️  python3 not found - couldn't test JSON response, but service is probably fine"
    echo "Manual test: curl http://localhost:9040/api/health | python3 -m json.tool"
fi

echo ""
echo "Installation complete!"
echo "Service status: systemctl status home-control-center"
echo "View logs: journalctl -u home-control-center -f"