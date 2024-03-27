package resources

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
)

func GetInstances(service *compute.Service, projectID string) ([]string, error) {
	// List instances aggregated by zone in the project
	instancesList, err := service.Instances.AggregatedList(projectID).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to list instances: %v", err)
	}

	var allInstances []*compute.Instance

	// Iterate over the zones and collect instances from each zone
	for _, zoneInstances := range instancesList.Items {
		allInstances = append(allInstances, zoneInstances.Instances...)
	}

	var instancelist []string
	for _, instance := range allInstances {
		jsonInstances, err := json.MarshalIndent(instance, "", " ")
		if err != nil {
			log.Fatalf("Error Marshalling instances to json: %v", err)
		}
		instancelist = append(instancelist, string(jsonInstances))
	}

	return instancelist, nil
}

func CreateInstance(service *compute.Service, projectID, instanceName, zone, machineType string) error {
	fmt.Printf("Creating instance %s in project %s\n", instanceName, projectID)

	// Create an instance resource object with the instance details
	instance := &compute.Instance{
		Name:        instanceName,
		MachineType: fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType),
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "projects/debian-cloud/global/images/family/debian-10",
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				AccessConfigs: []*compute.AccessConfig{
					{
						Name: "External NAT",
						Type: "ONE_TO_ONE_NAT",
					},
				},
				Network: "global/networks/default",
			},
		},
	}

	// Call the Instances.Insert method to create the instance
	op, err := service.Instances.Insert(projectID, zone, instance).Do()
	if err != nil {
		return fmt.Errorf("failed to create instance: %v", err)
	}

	waitVMOperation(service, projectID, zone, op.Name, "creating")
	fmt.Printf("Instance %s created successfully!\n", instanceName)
	return nil
}

func DeleteInstance(service *compute.Service, projectID, instanceName, zone string) error {
	fmt.Printf("Deleting instance %s in project %s\n", instanceName, projectID)

	// Call the Instances.Delete method to create the instance
	op, err := service.Instances.Delete(projectID, zone, instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to delete instance: %v", err)
	}

	waitVMOperation(service, projectID, zone, op.Name, "deleting")
	fmt.Printf("Instance %s deleted successfully!\n", instanceName)
	return nil
}

func ChangeInstanceState(service *compute.Service, projectID, zone, instanceName, state string) error {
	fmt.Printf("Attempting to %s instance %s in project %s...\n", state, instanceName, projectID)

	// Check the current instance state and perform the desired state change
	switch state {
	case "START":
		opr, err := service.Instances.Start(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitVMOperation(service, projectID, zone, opr.Name, state+"ing")

	case "STOP":
		opr, err := service.Instances.Stop(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitVMOperation(service, projectID, zone, opr.Name, state+"ing")
	case "RESTART":
		opr, err := service.Instances.Reset(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitVMOperation(service, projectID, zone, opr.Name, state+"ing")

	}

	return nil
}

func UpdateInstance() {

}

func SomeMultiplier(data int) {
	data *= data
	fmt.Println("Print the number square", data)
}
