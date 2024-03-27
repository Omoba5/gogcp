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

	doAction := 3

	ctx := context.Background()
	service, _ := newComputeService(ctx)

	projectID := os.Getenv("projectID")
	instanceNameTemplate := "test-instance-%d"
	zone := "us-central1-a"
	machineType := "e2-small"

	switch doAction {
	case 0:
		// Get Instance list
		instances, err := resources.GetInstances(service, projectID)
		if err != nil {
			log.Fatalf("Failed to get instances: %v", err)
		}

		// Print the details of each instance
		fmt.Println(instances[0])

	case 1:
		// Create Instances
		instanceName := instanceNameTemplate
		err := resources.CreateInstance(service, projectID, instanceName, zone, machineType)
		if err != nil {
			log.Fatalf("Failed to create instance: %v", err)
		}

	case 2:
		// Destroy Instances
		instanceName := "target-ssh"
		err := resources.DeleteInstance(service, projectID, instanceName, zone)
		if err != nil {
			log.Fatalf("Failed to delete instance: %v", err)
		}

	case 3:
		// Change Instance state
		instanceName := "cims-ssh"
		err := resources.ChangeInstanceState(service, projectID, zone, instanceName, "START")
		if err != nil {
			log.Fatalf("Failed to START instance: %v", err)
		}

	case 4:
		// Update Instance

	}

}
