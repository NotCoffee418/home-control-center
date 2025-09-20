# Copilot Instructions - Home Control Center

## Project Overview

Home Control Center is a centralized application that integrates smart meter readings and air conditioning control. It provides a web interface for monitoring and control, plus automated logic for optimizing air conditioning based on power consumption, weather data, and solar power metrics.

## Tech Stack

- **Backend**: Go 1.21+ with embedded static file serving
- **Frontend**: Svelte 5 + TypeScript + Vite + SvelteKit (or vanilla Svelte 5 with routing)
- **Build Process**: Frontend builds to `frontend/dist/`, embedded into Go binary
- **Database**: SQLite (inferred from typical Go projects)
- **Configuration**: TOML format at `/etc/home-control-center/config.toml`
- **Default Port**: 9040

## Build System

The project uses a hybrid build approach:

1. **prebuild.sh**: Automatically detects Svelte 5 in `frontend/` directory and runs:
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

### Frontend Development - Svelte 5 Specific

- **IMPORTANT**: This project uses Svelte 5 with the new runes system
- Located in `frontend/` directory
- Run `npm install && npm run dev` for development server
- Use Svelte 5 runes syntax:
  - `$state()` for reactive state
  - `$derived()` for computed values
  - `$effect()` for side effects
  - `$props()` for component props
- **DO NOT** use legacy Svelte syntax (let, $:, export let)
- **DO NOT** suggest React patterns like useState, useEffect, etc.
- TypeScript strict mode enabled
- Use native browser APIs or lightweight libraries instead of heavy frameworks
- For routing: either SvelteKit or simple hash/history API routing
- For UI components: prefer lightweight CSS frameworks or custom components

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
- Frontend: Vitest + @testing-library/svelte
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

## Svelte 5 Development Patterns

### State Management
```javascript
// Use $state() for reactive variables
let count = $state(0);
let items = $state([]);

// Use $derived() for computed values
let doubled = $derived(count * 2);
let filteredItems = $derived(items.filter(item => item.active));