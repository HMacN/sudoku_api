package example

import "sudoku_api/services/logging"

/*
This seems a bit ugly.  I don't like having a singleton to share the pointer to the Fx service with the Cobra
command, but I don't think I'll ever need to return more than one instance of the service.  For now, it works, and
is private to the package.  If I ever run into problems (probably while writing unit tests) I'll revisit this
architecture.
*/
var singleton *Service

func NewService(logger logging.LogWrapper) *Service {
	service := &Service{logger: logger}
	singleton = service
	return service
}

type Service struct {
	logger logging.LogWrapper
}

func (s *Service) Foo() {
	s.logger.LogInfo("Example")
}
