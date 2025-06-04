package initialization

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"isonetric-mmo-backend/pkg/model"
)

func Config() (*model.Config, error) {
	config := createDefaultConfig()

	if err := loadConfigData(config); err != nil {
		return nil, err
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func createDefaultConfig() *model.Config {
	return &model.Config{
		Server: &model.ServerConfig{
			Port: 8080,
		},
		Logging: &model.LoggingConfig{
			Disabled: false,
			Level:    "info",
		},
	}
}

func loadConfigData(config *model.Config) error {
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

func validateConfig(config *model.Config) error {
	validate := validator.New()

	if err := validate.Struct(config); err != nil {
		return err
	}

	return nil
}
