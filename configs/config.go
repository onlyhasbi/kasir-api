package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func LoadingConfig() (*Config, error) {
	var config Config

	_ = viper.BindEnv("DBConn", "DB_CONN")
	_ = viper.BindEnv("Port", "PORT")

	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
