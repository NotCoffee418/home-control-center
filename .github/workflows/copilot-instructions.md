# Copilot Instructions - Home Control Center

## Project Overview

Home Control Center is a centralized application that integrates smart meter readings and air conditioning control. It provides a web interface for monitoring and control, plus automated logic for optimizing air conditioning based on power consumption, weather data, and solar power metrics.

## Tech Stack

- **Backend**: Go 1.21+ with embedded static file serving
- **Frontend**: React + TypeScript + Vite + Mantine UI + React Router
- **Build Process**: Frontend builds to `frontend/dist/`, embedded into Go binary
- **Database**: SQLite (inferred from typical Go projects)
- **Configuration**: TOML format at `/etc/home-control-center/config.toml`
- **Default Port**: 9040

## Build System

The project uses a hybrid build approach:

1. **prebuild.sh**: Automatically detects React in `frontend/` directory and runs:
   - `npm install` (if node_modules missing)
   - `npm run build` to generate `frontend/dist/`
2. **Go Build**: Embeds `frontend/dist/` files using Go's embed functionality
   - Final binary serves both API and static frontend files
   - Single deployable artifact

## API Endpoints

- Health check available at `/api/health`
- API routes defined in `internal/web/api_routes.go`
- Additional endpoints to be documented as they're implemented

## External Service Integration

### ESP32 IR Airco Controller

- Communicates via HTTP API to ESP32 device
- Controls air conditioning units via infrared signals
- Expects JSON payloads for temperature, mode, and power commands

### European Smart Meter Integration

- Reads power consumption data from smart meters
- Likely uses serial/USB connection or network protocol
- Processes COSEM/DLMS or similar European smart meter standards

## Development Notes

### Frontend Development

- Located in `frontend/` directory
- Run `npm install && npm run dev` for development server
- Mantine provides UI components - prefer using existing components
- TypeScript strict mode enabled
- React Router handles client-side routing

### Backend Development

- Main server code in `internal/web/`
- Configuration handling in `internal/config/`
- Database operations in `internal/db/`
- Path utilities in `internal/pathing/`

### Key Features

- Real-time power consumption monitoring
- Automated air conditioning control
- Weather data integration
- Solar power monitoring
- Web dashboard for system monitoring

### Testing Approach

- Go: Standard `go test` framework
- Frontend: Jest + React Testing Library
- API: Health check at `/api/health` for basic connectivity
- Integration: Test external device communication carefully

### Configuration Management

- TOML format for human-readable config
- Environment variable overrides supported
- Default values should allow basic operation
- Sensitive data (API keys) should use environment variables

### Deployment Considerations

- Single binary deployment (frontend embedded)
- Systemd service file recommended
- Log rotation for long-running operation
- Graceful shutdown handling for device connections
- Network connectivity required for weather and external device APIs

## Error Handling Patterns

- Graceful degradation when external devices unavailable
- Retry logic for network-dependent operations
- User-friendly error messages in frontend
- Structured logging for debugging

## Security Notes

- Internal network deployment assumed
- HTTPS recommended if exposed externally
- Input validation on all API endpoints
