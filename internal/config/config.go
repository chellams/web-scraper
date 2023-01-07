package config

// ServerConfig holds the server configurations
type ServerConfig struct {
	IsGRPCEnabled bool   `cfg:"enable_grpc" cfgDefault:"false"`
	Address       string `cfg:"address" cfgDefault:"localhost:9876"`
	LogLevel      string `cfg:"log_level" cfgDefault:"info"`
}
