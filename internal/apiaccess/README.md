# API Access Module

This module provides access to external API endpoints for the Home Control Center application.

## Supported APIs

### AC Controller (ESP32 IR Controller)
- Control air conditioning units via HTTP API
- Endpoints: `/control`, `/status`, `/ping`
- Multiple devices supported via configuration

### Smart Meter API
- Read power consumption and production data
- European smart meter integration
- Endpoint: `/current`, `/ping`

## Response Structures

### AC Controller Response Structure
The `ACStatusResponse` struct is implemented based on typical ESP32 IR controller implementations, but the actual response format needs to be confirmed with the specific ESP32 device implementation.

**Expected JSON structure:**
```json
{
  "device_id": "string",
  "temperature": 25,
  "mode": "cool",
  "power": true,
  "fan_speed": "auto",
  "online": true
}
```

**Required confirmation:** 
- Exact field names from ESP32 implementation
- Available modes (cool, heat, auto, fan, dry, etc.)
- Fan speed options (low, medium, high, auto, etc.)
- Any additional fields like humidity, timer settings, etc.

### Smart Meter Response Structure
The `SmartMeterResponse` struct is based on typical European smart meter COSEM/DLMS data structures, but needs confirmation with the actual smart meter implementation.

**Expected JSON structure:**
```json
{
  "timestamp": "2024-01-01T12:00:00Z",
  "active_power_import_w": 1500.0,
  "active_power_export_w": 0.0,
  "energy_import_total_kwh": 1234.56,
  "energy_export_total_kwh": 78.90,
  "voltage_v": 230.0,
  "current_a": 6.5,
  "power_factor": 0.95,
  "online": true
}
```

**Required confirmation:**
- Exact field names from smart meter API
- Units of measurement (W, kW, kWh)
- Whether reactive power data is available
- Phase-specific data for 3-phase connections
- Timestamp format (ISO 8601, Unix timestamp, etc.)
- Any additional tariff or billing information

## Usage

```go
// Create API service
apiService := apiaccess.NewAPIService()

// Control AC device
err := apiService.ControlAC(ctx, "LivingRoom", apiaccess.ACControlRequest{
    Temperature: 22,
    Mode: "cool",
    Power: true,
})

// Get smart meter data
meterData, err := apiService.GetSmartMeterData(ctx)
```

## Configuration

The module uses endpoints defined in `config.toml`:

```toml
smart_meter_api_endpoint = "http://192.168.1.200/api/meter"

[ac_controller_endpoints]
"LivingRoom" = "http://192.168.1.100/api"
"Bedroom" = "http://192.168.1.101/api"
```

## Error Handling

All functions return detailed error messages with context about:
- Network connectivity issues
- HTTP status code errors
- JSON parsing errors
- Configuration issues (missing endpoints)
- Timeout errors (30-second default timeout)