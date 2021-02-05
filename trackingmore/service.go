package trackingmore

import "context"

// Service is
type Service interface {
	GetTracking(ctx context.Context, trackingNumber string) (string, error)
}
