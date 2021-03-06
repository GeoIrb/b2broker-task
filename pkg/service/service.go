package service

import (
	"context"
	"time"
)

// Service load
type Service struct{}

// Handler to requests
func (s *Service) Handler(ctx context.Context, requestUUID string) (err error) {
	var timeout time.Duration
	for _, s := range requestUUID {
		timeout += time.Duration(s)
	}
	time.Sleep(timeout * time.Millisecond)
	return
}

// New load service
func New() *Service {
	return &Service{}
}
