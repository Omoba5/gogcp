package resources

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

func CreateFirewallRule(service *compute.Service, projectID, firewallName string) error {
	fmt.Printf("Creating firewall rule %s in project %s\n", firewallName, projectID)

	// Create a firewall resource object with the firewall details
	firewall := &compute.Firewall{
		Name: firewallName,
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"80", "443"},
			},
		},
		SourceRanges: []string{"0.0.0.0/0"},
	}

	// Call the Firewalls.Insert method to create the firewall rule
	op, err := service.Firewalls.Insert(projectID, firewall).Do()
	if err != nil {
		return fmt.Errorf("failed to create firewall rule: %v", err)
	}

	waitFWOperation(service, projectID, op.Name, "creating")
	fmt.Printf("Firewall rule %s created successfully!\n", firewallName)
	return nil
}

func DeleteFirewallRule(service *compute.Service, projectID, firewallName string) error {
	fmt.Printf("Deleting firewall rule %s in project %s\n", firewallName, projectID)

	// Call the Firewalls.Delete method to delete the firewall rule
	op, err := service.Firewalls.Delete(projectID, firewallName).Do()
	if err != nil {
		return fmt.Errorf("failed to delete firewall rule: %v", err)
	}

	waitFWOperation(service, projectID, op.Name, "deleting")
	fmt.Printf("Firewall rule %s deleted successfully!\n", firewallName)
	return nil
}
