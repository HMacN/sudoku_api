package run_server

import (
	"net/http"
	"sudoku_api/services/logging"
)

var singleton *ServerService

func NewServerService(logger logging.LogWrapper, server *http.Server) *ServerService {
	singleton = &ServerService{
		logger: logger,
		server: server,
	}
	return singleton
}

type ServerService struct {
	logger logging.LogWrapper
	server *http.Server
}

func (s *ServerService) RunServer() error {
	s.logger.LogInfo("Starting server...")
	return s.server.ListenAndServe()
}
