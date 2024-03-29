package resources

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

// CreateNetwork creates a new network with the specified name in the specified project.
func CreateNetwork(service *compute.Service, projectID, networkName string) error {
	fmt.Printf("Creating network %s in project %s\n", networkName, projectID)

	subnet := &compute.Subnetwork{
		Name:        networkName,
		Network:     fmt.Sprintf("projects/%s/global/networks/%s", projectID, networkName),
		IpCidrRange: "10.138.0.0/20",
		Region:      "us-west1",
	}

	// Define the network resource
	network := &compute.Network{
		Name:                  networkName,
		AutoCreateSubnetworks: false,
		ForceSendFields:       []string{"AutoCreateSubnetworks"},
	}

	// Perform the network creation
	op, err := service.Networks.Insert(projectID, network).Do()
	if err != nil {
		return fmt.Errorf("failed to create network: %v", err)
	}

	// Wait for the operation to complete
	if err := waitGlobalOperation(service, projectID, op.Name, "creating network"); err != nil {
		return err
	}

	op, err2 := service.Subnetworks.Insert(projectID, "us-west1", subnet).Do()
	if err2 != nil {
		return fmt.Errorf("failed to create subnetwork: %v", err)
	}

	// Wait for the operation to complete
	if err := waitRegionOperation(service, projectID, "us-west1", op.Name, "creating network"); err != nil {
		return err
	}

	fmt.Printf("Network %s created successfully!\n", networkName)
	return nil
}

// DeleteNetwork deletes the network with the specified name from the specified project.
func DeleteNetwork(service *compute.Service, projectID, networkName string) error {
	fmt.Printf("Deleting network %s from project %s\n", networkName, projectID)

	// Perform the network deletion
	op, err := service.Networks.Delete(projectID, networkName).Do()
	if err != nil {
		return fmt.Errorf("failed to delete network: %v", err)
	}

	// Wait for the operation to complete
	if err := waitGlobalOperation(service, projectID, op.Name, "deleting network"); err != nil {
		return err
	}

	fmt.Printf("Network %s deleted successfully!\n", networkName)
	return nil
}
