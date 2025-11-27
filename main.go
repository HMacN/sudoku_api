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
	cmd.Execute()
}
