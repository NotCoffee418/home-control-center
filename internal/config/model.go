package config

type Config struct {
	DatabasePath  string `toml:"database_path"`
	ListenAddress string `toml:"listen_address"`
	ListenPort    int    `toml:"listen_port"`
}

