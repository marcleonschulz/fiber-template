package config

import (
	"github.com/spf13/viper"
)

var Config ViperConfig

func Init() error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	var err error
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Config.Db)
	err = viper.Unmarshal(&Config.Api)
	if err != nil {
		return err
	}
	return nil
}
