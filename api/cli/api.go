package cli

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/zeelrupapara/resumekraft/config"
	"github.com/zeelrupapara/resumekraft/database"
	"github.com/zeelrupapara/resumekraft/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// GetAPICommandDef runs app
func GetAPICommandDef(cfg config.AppConfig, logger *zap.Logger) cobra.Command {
	apiCommand := cobra.Command{
		Use:   "api",
		Short: "To start api",
		Long:  `To start api`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Create fiber app
			app := fiber.New(fiber.Config{})
			logger.Debug("Resumekraft server app created")

			// create database connection
			db, err := database.Connect(cfg.DB)
			if err != nil {
				logger.Panic("error while connecting to database", zap.Error(err))
				return err
			}
			logger.Debug("Database connection established")

			// setup routes
			err = routes.Setup(app, db, logger, cfg)
			if err != nil {
				logger.Panic("error while setting up routes", zap.Error(err))
				return err
			}
			logger.Debug("Routes are setup")

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				if err := app.Listen(cfg.Port); err != nil {
					logger.Panic(err.Error())
				}
			}()

			<-interrupt
			logger.Info("gracefully shutting down...")
			if err := app.Shutdown(); err != nil {
				logger.Panic("error while shutdown server", zap.Error(err))
			}

			logger.Info("server stopped to receive new requests or connection.")
			return nil
		},
	}

	return apiCommand
}
