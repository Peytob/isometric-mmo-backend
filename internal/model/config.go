package model

type Config struct {
	Server  *ServerConfig
	Logging *LoggingConfig
}

type ServerConfig struct {
	Port int `validate:"required,min=1,max=65535"`
}

type LoggingConfig struct {
	Disabled bool
	Level    string `validate:"required,oneof=debug info warn error"`
}
