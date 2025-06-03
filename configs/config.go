package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int `validate:"required,min=1,max=65535"`
}

func LoadConfig() (*Config, error) {
	config := createDefaultConfig()

	if err := loadConfigData(config); err != nil {
		return nil, err
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func createDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: 8080,
		},
	}
}

func loadConfigData(config *Config) error {
	// Probably too many overhead from viper lib. I just need to load configuration from file and envs, nothing more

	viper.AddConfigPath("./")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

func validateConfig(config *Config) error {
	validate := validator.New()

	if err := validate.Struct(config); err != nil {
		return err
	}

	return nil
}
