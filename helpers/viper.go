package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

func Env(key string) string {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error getting key from .env: ", err.Error())
		panic("viper error")
	}

	return viper.GetString(key)
}
