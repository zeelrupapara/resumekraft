// Golang API.
//
//	Schemes: https
//	Host: localhost
//	BasePath: /api/v1
//	Version: 0.0.1-alpha
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"github.com/zeelrupapara/resumekraft/cli"
	"github.com/zeelrupapara/resumekraft/config"
	"github.com/zeelrupapara/resumekraft/logger"
	"go.uber.org/zap"
)

func main() {
	// Collecting config from env or file or flag
	cfg := config.GetConfig()

	logger, err := logger.NewRootLogger(cfg.Debug, cfg.IsDevelopment)
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	err = cli.Init(cfg, logger)
	if err != nil {
		panic(err)
	}

}
