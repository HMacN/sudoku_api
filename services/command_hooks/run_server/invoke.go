package run_server

import (
	"net/http"
	"sudoku_api/services/logging"
)

func invoke(logger logging.LogWrapper, server *http.Server) {
	logger.LogInfo("Starting server...")
	err := server.ListenAndServe()
	if err != nil {
		logger.LogError("Error starting server: %v", err)
	}
}
