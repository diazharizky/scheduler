//go:generate go run -v github.com/go-swagger/go-swagger/cmd/swagger generate spec -m -o ../../api/swagger-spec/scheduler.json
//go:generate go run -v github.com/gobuffalo/packr/v2/packr2

// Package main Scheduler.
//
// The purpose of this application is to provice service that allows you to run specific routine/cron at scheduled time, either it's once or forever
//
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: scheduler:3000
//     BasePath: /
//     Version: 0.0.0
//     Contact: Diaz Harizky<diazharizky@gmail.com>
//
// swagger:meta
package main

import (
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
