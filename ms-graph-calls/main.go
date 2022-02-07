package main

import (
	"fmt"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/spf13/viper"
	graph "msgraphcalls/graph"
)

func main(){
	
	var msGraphOptions graph.MSGraphOptions
	err := viper.Sub("msGraphOptions").Unmarshal(&msGraphOptions)
	if err != nil {
		fmt.Printf("[ERR] Failed to unmarshal SampleOption: %v", err)
	}
}

func loadConfig() {
	env := "dev"
	
	// - Setting defaults
	viper.SetDefault("TestConfigKey", "DefaultValue")

	// - Loading config from file
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath("./configurations")
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