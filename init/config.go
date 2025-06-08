package init

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"isonetric-mmo-backend/internal/model"
	"strings"
)

func Config() (*model.Config, error) {
	config := createDefaultConfig()

	if err := loadConfigData(config); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if err := validateConfig(config); err != nil {
		// todo make more readability text via validator.ValidationErrors
		return nil, fmt.Errorf("config valudation failed. Details: %w", err)
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
		Store: &model.StoreConfig{
			Sql: &model.SqlConfig{
				Host:     "localhost",
				Username: "postgres",
				Password: "",
				Port:     5432,
				DbName:   "mmobackend",
			},
		},
	}
}

func loadConfigData(config *model.Config) error {
	// Probably too many overhead from viper lib. I just need to load configuration from file and envs, nothing more
	// and there are some problems with env values deserializing. Used some hack from
	// https://github.com/spf13/viper/issues/188

	if err := setupDefaultViperConfig(config); err != nil {
		return err
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("mmo")
	viper.AutomaticEnv()

	viper.AddConfigPath("./")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

func setupDefaultViperConfig(config *model.Config) error {
	viper.SetConfigType("yaml")

	defaultConfigBinary, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	defaultConfig := bytes.NewReader(defaultConfigBinary)
	if err := viper.MergeConfig(defaultConfig); err != nil {
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
