package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	
	loadConfig()
	
	fmt.Println(viper.GetString("TestConfigKey"))
}

func loadConfig() {
	env := "dev"
	
	// - Setting defaults
	viper.SetDefault("TestConfigKey", "DefaultValue")

	// - Loading config from file
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[WRN] No config files found")
		} else {
			panic(fmt.Errorf("[ERR] config failed to load: %v", err))
		}
	}

	// - Loading config from env
	viper.SetEnvPrefix("BLT")
	viper.AutomaticEnv()
}