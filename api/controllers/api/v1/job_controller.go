package v1

import (
	"github.com/zeelrupapara/resumekraft/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"github.com/doug-martin/goqu/v9"
)

type JobController struct {
	jobModel *models.JobModel
	logger   *zap.Logger
}

// NewJobController returns a user
func NewJobController(db *goqu.Database, logger *zap.Logger) (*JobController, error) {
	jobModel, err := models.InitJobModel(db)
	if err != nil {
		logger.Error("error while initializing job model", zap.Error(err))
		return nil, err
	}
	return &JobController{
		jobModel: jobModel,
		logger:   logger,
	}, nil
}

// SendJob to workers and process job
func (ctrl *JobController) SendJob(c *fiber.Ctx) error {
	return nil
}