package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gogcp/resources"

	"github.com/joho/godotenv"
	"google.golang.org/api/compute/v1"
)

func newComputeService(ctx context.Context) (*compute.Service, error) {
	// Create a new Compute Engine service client
	service, err := compute.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %v", err)
	}
	return service, nil
}

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()
	service, _ := newComputeService(ctx)

	projectID := os.Getenv("projectID")
	instanceNameTemplate := "test-instance-%d"
	zone := "us-central1-a"
	// machineType := "e2-small"

	// //We can use go routines to create instances in parallel
	// for i := 0; i < 2; i++ {
	// 	instanceName := fmt.Sprintf(instanceNameTemplate, i)
	// 	err := resources.CreateInstance(service, projectID, instanceName, zone, machineType)
	// 	if err != nil {
	// 		log.Fatalf("Failed to create instance: %v", err)
	// 	}
	// }

	//We can use go routines to create instances in parallel
	for i := 0; i < 2; i++ {
		instanceName := fmt.Sprintf(instanceNameTemplate, i)
		err := resources.DeleteInstance(service, projectID, instanceName, zone)
		if err != nil {
			log.Fatalf("Failed to delete instance: %v", err)
		}
	}

	resources.SomeMultiplier(23)
}
