package utils

import "github.com/spf13/viper"

type Config struct {
	DBUrl string `mapstructure:"DATABASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if viper.ReadInConfig() != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
