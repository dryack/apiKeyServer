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

		// fmt.Println("Config file changed: ", e.Name)
		tempKeys := Keys{
			TotalExhaustions: keys.TotalExhaustions,
			TotalKeysServed:  keys.TotalKeysServed,
			ServerVersion:    keys.ServerVersion,
			StartupTime:      keys.StartupTime,
		}
		if err := viper.Unmarshal(&tempKeys); err != nil {
			fmt.Println("fucking fuck")
			Log.Warn().Msg("new configuration failed to unmarshal")
		} else {
			Log.Info().Msg("new configuration loaded")
			liveUpdateConfiguration(&tempKeys)
		}
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

func liveUpdateConfiguration(tempKeys *Keys) {
	for _, nKey := range tempKeys.Apikeys {
		for _, oKey := range keys.Apikeys {
			if nKey.User == oKey.User {
				if int32(nKey.MaxPerMinute) < oKey.CurrentlyRemaining {
					nKey.CurrentlyRemaining = int32(nKey.MaxPerMinute)
				} else {
					nKey.CurrentlyRemaining = oKey.CurrentlyRemaining
				}
				nKey.Active = oKey.Active
				nKey.Tornkey = oKey.Tornkey
				nKey.ReturnToService = oKey.ReturnToService
				nKey.Kills = oKey.Kills
				nKey.Uses = oKey.Uses
			} else {
				nKey.CurrentlyRemaining = int32(nKey.MaxPerMinute)
			}

		}
	}
	for _, key := range tempKeys.Apikeys {
		tempKeys.TotalPerMinute += key.MaxPerMinute
	}
	fmt.Println(tempKeys)
	// newConf := viper.AllSettings()
}
