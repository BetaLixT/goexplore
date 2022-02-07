package main

import (
	"fmt"
	"os"

	graph "msgraphcalls/graph"
	graphHelper "msgraphcalls/graph"
	"github.com/spf13/viper"
)

func main(){
	
	loadConfig()

	var msGraphOptions graph.MSGraphOptions
	msGraphOptionsSub := viper.Sub("msGraphOptions")
	if msGraphOptionsSub == nil {
		fmt.Printf("[ERR] failed to fetch required options\n")
		os.Exit(1)
	}

	err := msGraphOptionsSub.Unmarshal(&msGraphOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("[ERR] Failed to unmarshal SampleOption: %v", err)
		os.Exit(1)
	}
	graphHelper, err := graphHelper.NewGraphHelper(msGraphOptions)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("[ERR] Failed to create graph helper")
		os.Exit(1)
	}
	user, err := graphHelper.GetUser("alphin@cxunicorn.com")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("[ERR] Failed to fetch users")
		os.Exit(1)
	}
	fmt.Printf("User: %s\n", *user.GetDisplayName())
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