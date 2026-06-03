package example

import (
	"sudoku_api/config"
	"sudoku_api/config/config_keys"
	"sudoku_api/services/logging"
)

func invoke(logger logging.LogWrapper) {
	exampleOne := "exampleOne:" + config.GetString(config_keys.KeyExampleOne)
	exampleTwo := "exampleTwo:" + config.GetString(config_keys.KeyExampleTwo)
	exampleThree := "exampleThree:" + config.GetString(config_keys.KeyExampleThree)
	exampleFour := "exampleFour:" + config.GetString(config_keys.KeyExampleFour)
	logger.LogInfo(exampleThree)
	logger.LogInfo(exampleFour)
	logger.LogInfo(exampleOne)
	logger.LogInfo(exampleTwo)
}
