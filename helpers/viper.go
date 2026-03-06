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
	viper.AutomaticEnv() 

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			val := viper.GetString(key)
			if val == "" {
				fmt.Printf("Warning: .env file not found and key '%s' is not set in environment\n", key)
			}
			return val
		}
		fmt.Println("error getting key from .env: ", err.Error())
		panic("viper error")
	}

	return viper.GetString(key)
}
