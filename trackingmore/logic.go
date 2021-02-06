package trackingmore

import (
	"context"
	"io/ioutil"
	"net/http"

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
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		logger.Log(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb, nil
}
