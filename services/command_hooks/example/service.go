package example

import (
	"sudoku_api/config"
	"sudoku_api/config/config_keys"
	"sudoku_api/services/logging"
)

/*
This seems a bit ugly.  I don't like having a singleton to share the pointer to the Fx service with the Cobra
command, but I don't think I'll ever need to return more than one instance of the service.  For now, it works, and
is private to the package.  If I ever run into problems (probably while writing unit tests) I'll revisit this
architecture.
*/
var singleton *Service

func NewService(logger logging.LogWrapper) *Service {
	singleton = &Service{logger: logger}
	logger.LogInfo("service initialised")
	return singleton
}

type Service struct {
	logger logging.LogWrapper
}

func (s *Service) Foo() {
	s.logger.LogInfo("Example")
	exampleOne := config.GetString(config_keys.KeyExampleOne)
	exampleTwo := config.GetString(config_keys.KeyExampleTwo)
	exampleThree := config.GetString(config_keys.KeyExampleThree)
	exampleFour := config.GetString(config_keys.KeyExampleFour)
	s.logger.LogInfo(exampleOne)
	s.logger.LogInfo(exampleTwo)
	s.logger.LogInfo(exampleThree)
	s.logger.LogInfo(exampleFour)
}
