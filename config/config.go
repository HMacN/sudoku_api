package config

import (
	"sudoku_api/config/config_keys"

	"github.com/spf13/viper"
)

func GetString(key config_keys.Key) string {
	return viper.GetString(key.String())
}

func GetInt(key config_keys.Key) int {
	return viper.GetInt(key.String())
}

func GetFloat(key config_keys.Key) float64 {
	return viper.GetFloat64(key.String())
}

func GetBool(key config_keys.Key) bool {
	return viper.GetBool(key.String())
}
