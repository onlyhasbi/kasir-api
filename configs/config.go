package configs

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func LoadingConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.Port == "" {
		config.Port = os.Getenv("PORT")
	}

	if config.DBConn == "" {
		config.DBConn = os.Getenv("DB_CONN")
	}

	return &config, nil
}
