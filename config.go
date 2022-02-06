package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func doConfiguration() error {
	// configuration handling
	//TODO: probably change to apikeyserver.config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("/config/")
	//TODO: log the location of the active configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//TODO:  needs to be logged
			_ = fmt.Errorf("config file not found")
			return err
		} else {
			_ = fmt.Errorf("fatal error with config file: %w\n", err)
			return err
		}
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		//TODO: will need to be sending to logfile
		fmt.Println("Config file changed: ", e.Name)
	})
	viper.WatchConfig()
	err := viper.Unmarshal(&keys)
	if err != nil {
		//TODO: needs to be logged
		_ = fmt.Errorf("unable to decode into struct, %v", err)
		return err
	}
	return err
}
