package trackingmore

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeTrackingReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetTrackingRequest
	vars := mux.Vars(r)

	req = GetTrackingRequest{
		ID: vars["id"],
	}
	return req, nil
}

type (
	// GetTrackingResponse is
	GetTrackingResponse struct {
		Tracking string `json:"tracking"`
	}

	// GetTrackingRequest is
	GetTrackingRequest struct {
		ID string `json:"id"`
	}
)
