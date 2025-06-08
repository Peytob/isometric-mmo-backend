package model

type Config struct {
	Server  *ServerConfig
	Logging *LoggingConfig
	Store   *StoreConfig
}

type ServerConfig struct {
	Port int `validate:"required,min=1,max=65535"`
}

type LoggingConfig struct {
	Disabled bool
	Level    string `validate:"required,oneof=debug info warn error"`
}

type StoreConfig struct {
	Sql *SqlConfig
}

type SqlConfig struct {
	Host     string `validate:"required,omitempty"`
	Username string `validate:"required,omitempty,max=63"`
	Password string `validate:"required,omitempty"`
	Port     int    `validate:"required,min=1,max=65535"`
	DbName   string `validate:"required,max=63"`
}
