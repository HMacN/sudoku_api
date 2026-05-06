package config

import "github.com/spf13/viper"

func Get(key Key) any {
	return viper.Get(string(key))
}
