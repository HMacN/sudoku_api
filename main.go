// Package main
//
// Documentation for my go project
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host:
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	Security:
//	- basic
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//
// swagger:meta
package main

import (
	"sudoku_api/cmd"
)

func main() {
	// TODO: Move Fx into here?
	// TODO: Cobra refactor as per https://cobra.dev/docs/tutorials/12-factor-app/

	cmd.Execute()
}
