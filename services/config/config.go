package config

import (
	"sudoku_api/services/config/config_keys"

	"github.com/spf13/viper"
)

func Get[T any](key config_keys.Key) T {
	return T(viper.Get(key.String()))
}
