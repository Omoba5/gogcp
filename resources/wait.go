package resources

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/compute/v1"
)

func waitZoneOperation(service *compute.Service, projectID, zone, operationName string, action string) error {
	for {
		// Check operation status
		op, err := service.ZoneOperations.Get(projectID, zone, operationName).Do()
		if err != nil {
			log.Fatalf("Failed to get operation: %v", err)
		}
		if op.Status == "DONE" {
			break
		}
		fmt.Printf("Waiting more 10 secs for %s operation\n", action)
		time.Sleep(10 * time.Second)

		fmt.Println("DONE!")
	}
	return nil
}

func waitGlobalOperation(service *compute.Service, projectID, operationName string, action string) error {
	for {
		// Check operation status
		op, err := service.GlobalOperations.Get(projectID, operationName).Do()
		if err != nil {
			log.Fatalf("Failed to get operation: %v", err)
		}
		if op.Status == "DONE" {
			break
		}
		fmt.Printf("Waiting more 10 secs for %s operation\n", action)
		time.Sleep(10 * time.Second)

		fmt.Println("DONE!")
	}
	return nil
}
