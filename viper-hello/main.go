package main

import (
	"fmt"
	"github.com/spf13/viper"
	"viperhello/options"
)

func main() {
	
	loadConfig()
	
	fmt.Println(viper.GetString("TestConfigKey"))

	var s options.SampleOption
	err := viper.Sub("SampleOption").Unmarshal(&s)
	if err != nil {
		fmt.Printf("[ERR] Failed to unmarshal SampleOption: %v", err)
	} else {
		fmt.Printf("Sample options: %v %v\n", s.ConfigInt, s.ConfigString)
	}
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