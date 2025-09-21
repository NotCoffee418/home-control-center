package apiaccess

import (
	"context"
	"log"
	"time"
)

// ExampleUsage demonstrates how to use the API access module
// This is for documentation purposes and won't be included in the binary
func ExampleUsage() {
	// Create API service with default timeout (30 seconds)
	apiService := NewAPIService()

	// Create API service with custom timeout
	customAPIService := NewAPIServiceWithTimeout(60 * time.Second)

	ctx := context.Background()

	// Example 1: Control an AC device
	log.Println("Controlling AC device...")
	err := apiService.ControlAC(ctx, "LivingRoom", ACControlRequest{
		Temperature: 22,
		Mode:        "cool",
		Power:       true,
		FanSpeed:    "auto",
	})
	if err != nil {
		log.Printf("Failed to control AC: %v", err)
	} else {
		log.Println("AC control successful")
	}

	// Example 2: Get AC device status
	log.Println("Getting AC device status...")
	status, err := apiService.GetACStatus(ctx, "LivingRoom")
	if err != nil {
		log.Printf("Failed to get AC status: %v", err)
	} else {
		log.Printf("AC Status: Power=%t, Temp=%d°C, Mode=%s, Online=%t",
			status.Power, status.Temperature, status.Mode, status.Online)
	}

	// Example 3: Get all configured AC devices
	devices := apiService.GetAllACDevices()
	log.Printf("Configured AC devices: %v", devices)

	// Example 4: Ping AC device to check connectivity
	log.Println("Pinging AC device...")
	err = apiService.PingACDevice(ctx, "LivingRoom")
	if err != nil {
		log.Printf("AC device ping failed: %v", err)
	} else {
		log.Println("AC device is reachable")
	}

	// Example 5: Get smart meter data
	log.Println("Getting smart meter data...")
	meterData, err := apiService.GetSmartMeterData(ctx)
	if err != nil {
		log.Printf("Failed to get smart meter data: %v", err)
	} else {
		log.Printf("Smart Meter: Import=%.2fW, Export=%.2fW, Voltage=%.1fV, Online=%t",
			meterData.ActivePowerImport, meterData.ActivePowerExport,
			meterData.Voltage, meterData.Online)
	}

	// Example 6: Ping smart meter
	log.Println("Pinging smart meter...")
	err = apiService.PingSmartMeter(ctx)
	if err != nil {
		log.Printf("Smart meter ping failed: %v", err)
	} else {
		log.Println("Smart meter is reachable")
	}

	// Example 7: Using custom timeout service for operations that might take longer
	log.Println("Using custom timeout service...")
	err = customAPIService.ControlAC(ctx, "LivingRoom", ACControlRequest{
		Temperature: 24,
		Power:       false, // Turn off
	})
	if err != nil {
		log.Printf("Failed to turn off AC: %v", err)
	} else {
		log.Println("AC turned off successfully")
	}
}