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

	doAction := 2

	ctx := context.Background()
	service, _ := newComputeService(ctx)

	projectID := os.Getenv("projectID")
	instanceNameTemplate := "test-instance-1"
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
		username := "coolname"
		password := "bigsecret"
		err := resources.CreateInstance(service, projectID, instanceName, zone, machineType, username, password)
		if err != nil {
			log.Fatalf("Failed to create instance: %v", err)
		}

	case 2:
		// Destroy Instances
		instanceName := instanceNameTemplate
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
		// Update Instance Disk
		instanceName := "cims-ssh"
		err := resources.UpdateBootDiskSize(service, projectID, zone, instanceName, 21)
		if err != nil {
			log.Fatalf("Failed to update instance: %v", err)
		}

	case 5:
		// Update Instance Machine Type
		instanceName := "cims-ssh"

		// First stop the instance
		err := resources.ChangeInstanceState(service, projectID, zone, instanceName, "STOP")
		if err != nil {
			log.Fatalf("Failed to STOP instance: %v", err)
		}

		// Change the instance tyoe
		err2 := resources.UpdateMachineType(service, projectID, zone, instanceName, "e2-small")
		if err2 != nil {
			log.Fatalf("Failed to START instance: %v", err)
		}

		// Start the instance after modifying it.
		err3 := resources.ChangeInstanceState(service, projectID, zone, instanceName, "START")
		if err3 != nil {
			log.Fatalf("Failed to START instance: %v", err)
		}

	case 6:
		// Modify Network Tags
		instanceName := "cims-ssh"
		networkTags := []string{"http-server", "something-nice"}
		err := resources.UpdateNetworkTags(service, projectID, zone, instanceName, networkTags)
		if err != nil {
			log.Fatalf("Failed to modify instance's Network Tags: %v", err)
		}

	case 7:
		// Create Firewall Rule
		firewallName := "veryhot"
		firewallPortMap := map[string][]string{
			"tcp": {"80", "8080", "443"},
		}
		firewallTargets := []string{"http-server"}

		err := resources.CreateFirewallRule(service, projectID, firewallName, firewallPortMap, firewallTargets)
		if err != nil {
			log.Fatalf("Failed to create firewall rule: %v", err)
		}

	case 8:
		// Modify Firewall Target tags
		firewallName := "veryhot"
		firewallTargets := []string{"http-server", "someexperiment"}

		err := resources.UpdateFirewallTargetTags(service, projectID, firewallName, firewallTargets)
		if err != nil {
			log.Fatalf("Failed to update firewall rule: %v", err)
		}

	case 9:
		// Modify Firewall Ports
		firewallName := "veryhot"
		firewallPortMap := map[string][]string{
			"tcp": {"80", "8080", "443"},
			//		"udp": {"90"},
		}

		err := resources.UpdateFirewallPorts(service, projectID, firewallName, firewallPortMap)
		if err != nil {
			log.Fatalf("Failed to update firewall rule: %v", err)
		}

	case 10:
		// Modify Firewall Source Ranges
		firewallName := "veryhot"
		firewallSources := []string{"192.7.0.0/24", "196.55.5.78"}

		err := resources.UpdateFirewallSourceRanges(service, projectID, firewallName, firewallSources)
		if err != nil {
			log.Fatalf("Failed to update firewall rule: %v", err)
		}

	case 11:
		// Delete Firewall Rule

		firewallName := "veryhot"

		err := resources.DeleteFirewallRule(service, projectID, firewallName)
		if err != nil {
			log.Fatalf("Failed to delete firewall rule: %v", err)
		}

	case 12:
		// Create Network
		networkName := "coolname2"

		err := resources.CreateNetwork(service, projectID, networkName)
		if err != nil {
			log.Fatalf("Failed to Create network: %v", err)
		}

	case 13:
		// Delete Network
		networkName := "coolname2"

		err := resources.DeleteNetwork(service, projectID, networkName)
		if err != nil {
			log.Fatalf("Failed to Create network: %v", err)
		}
	}

}
