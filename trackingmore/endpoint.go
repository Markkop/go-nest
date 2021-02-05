package trackingmore

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is
type Endpoints struct {
	GetTracking endpoint.Endpoint
}

// MakeEndpoints is
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetTracking: makeGetTrackingEndpoint(s),
	}
}

func makeGetTrackingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTrackingRequest)
		tracking, err := s.GetTracking(ctx, req.ID)
		return GetTrackingResponse{
			Tracking: tracking,
		}, err
	}
}
