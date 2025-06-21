#!/bin/bash
# build.sh - Build frontend then Go binary

set -e  # Exit on any error

echo "Building frontend..."
cd frontend

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    echo "Installing npm dependencies..."
    npm install
fi

# Build with verbose output
echo "Running npm run build..."
npm run build

# Verify build output
if [ ! -d "dist" ]; then
    echo "ERROR: dist directory not created"
    exit 1
fi

if [ ! -f "dist/index.html" ]; then
    echo "ERROR: index.html not found in dist"
    exit 1
fi

echo "Frontend build successful"
ls -la dist/