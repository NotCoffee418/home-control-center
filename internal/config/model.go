package config

type Config struct {
	DatabasePath            string            `toml:"database_path"`
	ListenAddress           string            `toml:"listen_address"`
	ListenPort              int               `toml:"listen_port"`
	AcControllerEndpoints   map[string]string `toml:"ac_controller_endpoints"`
	SmartMeterApiEndpoint   string            `toml:"smart_meter_api_endpoint"`
}

