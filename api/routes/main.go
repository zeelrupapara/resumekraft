package routes

import (
	"sync"

	"go.uber.org/zap"

	"github.com/zeelrupapara/resumekraft/config"
	controller "github.com/zeelrupapara/resumekraft/controllers/api/v1"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

// Setup func
func Setup(app *fiber.App, goqu *goqu.Database, logger *zap.Logger, config config.AppConfig) error {
	mu.Lock()

	router := app.Group("/api")
	v1 := router.Group("/v1")

	err := setupJobController(v1, goqu, logger)
	if err != nil {
		return err
	}

	mu.Unlock()
	return nil
}

// Sending jobs description
func setupJobController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger) error {
	jobController, err := controller.NewJobController(goqu, logger)
	if err != nil {
		return err
	}

	v1.Post("/job", jobController.SendJob)

	return nil
}

