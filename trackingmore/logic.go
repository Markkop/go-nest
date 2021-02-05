package trackingmore

import (
	"context"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
}

// NewService is
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetTracking(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")
	logger.Log("got tracking")
	return id, nil
}
