package app

import (
	"stokeria/helper"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	return config
}