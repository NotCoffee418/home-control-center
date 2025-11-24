package apiaccess

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NotCoffee418/home-control-center/internal/config"
)

// ACMode represents the AC operation mode
type ACMode string

const (
	ACModeOff  ACMode = "off"
	ACModeCool ACMode = "cool"
	ACModeHeat ACMode = "heat"
	ACModeAuto ACMode = "auto"
	ACModeFan  ACMode = "fan"
	ACModeDry  ACMode = "dry"
)

// ACFanSpeed represents the AC fan speed setting
type ACFanSpeed string

const (
	ACFanSpeedAuto   ACFanSpeed = "auto"
	ACFanSpeedLow    ACFanSpeed = "low"
	ACFanSpeedMedium ACFanSpeed = "medium"
	ACFanSpeedHigh   ACFanSpeed = "high"
)

// APIService provides access to external API endpoints
type APIService struct {
	client HTTPClient
	config *config.Config
}

// NewAPIService creates a new API service with default timeout
func NewAPIService() *APIService {
	return &APIService{
		client: NewHTTPClient(30 * time.Second), // 30 second timeout
		config: config.GetConfig(),
	}
}

// ACControlRequest represents a request to control an air conditioning unit
type ACControlRequest struct {
	Temperature int        `json:"temperature,omitempty"`
	Mode        ACMode     `json:"mode,omitempty"`
	Power       bool       `json:"power"`
	FanSpeed    ACFanSpeed `json:"fan_speed,omitempty"`
}

// ACStatusResponse represents the response from AC controller status endpoint
type ACStatusResponse struct {
	DeviceID    string     `json:"device_id"`
	Temperature int        `json:"temperature"`
	Mode        ACMode     `json:"mode"`
	Power       bool       `json:"power"`
	FanSpeed    ACFanSpeed `json:"fan_speed"`
	Online      bool       `json:"online"`
}

// SmartMeterResponse represents the response from smart meter API
// Note: Actual response structure needs to be confirmed with smart meter implementation
// This is based on typical European smart meter COSEM/DLMS data
type SmartMeterResponse struct {
	Timestamp           string  `json:"timestamp"`
	ActivePowerImport   float64 `json:"active_power_import_w"`   // Current power consumption in watts
	ActivePowerExport   float64 `json:"active_power_export_w"`   // Current power export in watts (solar)
	EnergyImportTotal   float64 `json:"energy_import_total_kwh"` // Total imported energy in kWh
	EnergyExportTotal   float64 `json:"energy_export_total_kwh"` // Total exported energy in kWh
	Voltage             float64 `json:"voltage_v"`               // Voltage in volts
	Current             float64 `json:"current_a"`               // Current in amperes
	PowerFactor         float64 `json:"power_factor"`            // Power factor
	Online              bool    `json:"online"`
}

// ControlAC sends a control command to the specified AC controller
// Endpoint: POST {endpoint}/control
// Expected ESP32 endpoint structure needs confirmation
func (s *APIService) ControlAC(ctx context.Context, deviceKey string, request ACControlRequest) error {
	endpoint, exists := s.config.AcControllerEndpoints[deviceKey]
	if !exists {
		return fmt.Errorf("AC controller endpoint not found for device: %s", deviceKey)
	}

	url := fmt.Sprintf("%s/control", endpoint)
	
	resp, err := s.client.Post(ctx, url, request)
	if err != nil {
		return fmt.Errorf("failed to control AC device %s: %w", deviceKey, err)
	}
	defer resp.Body.Close()

	if err := checkHTTPStatus(resp); err != nil {
		return fmt.Errorf("AC control request failed for device %s: %w", deviceKey, err)
	}

	return nil
}

// GetACStatus retrieves the current status of the specified AC controller
// Endpoint: GET {endpoint}/status
// Response structure needs confirmation with ESP32 implementation
func (s *APIService) GetACStatus(ctx context.Context, deviceKey string) (*ACStatusResponse, error) {
	endpoint, exists := s.config.AcControllerEndpoints[deviceKey]
	if !exists {
		return nil, fmt.Errorf("AC controller endpoint not found for device: %s", deviceKey)
	}

	url := fmt.Sprintf("%s/status", endpoint)
	
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get AC status for device %s: %w", deviceKey, err)
	}
	defer resp.Body.Close()

	if err := checkHTTPStatus(resp); err != nil {
		return nil, fmt.Errorf("AC status request failed for device %s: %w", deviceKey, err)
	}

	body, err := readResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read AC status response for device %s: %w", deviceKey, err)
	}

	var status ACStatusResponse
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("failed to parse AC status response for device %s: %w", deviceKey, err)
	}

	return &status, nil
}

// GetAllACDevices returns a list of all configured AC device keys
func (s *APIService) GetAllACDevices() []string {
	devices := make([]string, 0, len(s.config.AcControllerEndpoints))
	for deviceKey := range s.config.AcControllerEndpoints {
		devices = append(devices, deviceKey)
	}
	return devices
}

// GetSmartMeterData retrieves current smart meter readings
// Endpoint: GET {endpoint}/current
// Response structure based on European smart meter standards, needs confirmation
func (s *APIService) GetSmartMeterData(ctx context.Context) (*SmartMeterResponse, error) {
	if s.config.SmartMeterApiEndpoint == "" {
		return nil, fmt.Errorf("smart meter API endpoint not configured")
	}

	url := fmt.Sprintf("%s/current", s.config.SmartMeterApiEndpoint)
	
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get smart meter data: %w", err)
	}
	defer resp.Body.Close()

	if err := checkHTTPStatus(resp); err != nil {
		return nil, fmt.Errorf("smart meter request failed: %w", err)
	}

	body, err := readResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read smart meter response: %w", err)
	}

	var meterData SmartMeterResponse
	if err := json.Unmarshal(body, &meterData); err != nil {
		return nil, fmt.Errorf("failed to parse smart meter response: %w", err)
	}

	return &meterData, nil
}

// PingACDevice checks if an AC controller device is reachable
// Endpoint: GET {endpoint}/ping
func (s *APIService) PingACDevice(ctx context.Context, deviceKey string) error {
	endpoint, exists := s.config.AcControllerEndpoints[deviceKey]
	if !exists {
		return fmt.Errorf("AC controller endpoint not found for device: %s", deviceKey)
	}

	url := fmt.Sprintf("%s/ping", endpoint)
	
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to ping AC device %s: %w", deviceKey, err)
	}
	defer resp.Body.Close()

	if err := checkHTTPStatus(resp); err != nil {
		return fmt.Errorf("AC device %s ping failed: %w", deviceKey, err)
	}

	return nil
}

// PingSmartMeter checks if the smart meter API is reachable
// Endpoint: GET {endpoint}/ping
func (s *APIService) PingSmartMeter(ctx context.Context) error {
	if s.config.SmartMeterApiEndpoint == "" {
		return fmt.Errorf("smart meter API endpoint not configured")
	}

	url := fmt.Sprintf("%s/ping", s.config.SmartMeterApiEndpoint)
	
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to ping smart meter: %w", err)
	}
	defer resp.Body.Close()

	if err := checkHTTPStatus(resp); err != nil {
		return fmt.Errorf("smart meter ping failed: %w", err)
	}

	return nil
}