package main

import (
	"context"
	"fmt"

	"google.golang.org/api/compute/v1"
)

func NewComputeService(ctx context.Context) (*compute.Service, error) {
	// Create a new Compute Engine service client
	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %v", err)
	}
	return service, nil
}
