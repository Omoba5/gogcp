package resources

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

func CreateInstance(service *compute.Service, projectID, instanceName, zone, machineType string) error {
	fmt.Printf("Creating instance %s in project %s\n", instanceName, projectID)

	// // Create a new Compute Engine service client
	// service, err := compute.NewService(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to create service: %v", err)
	// }

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

	// // Create a new Compute Engine service client
	// service, err := compute.NewService(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to create service: %v", err)
	// }

	// Call the Instances.Delete method to create the instance
	op, err := service.Instances.Delete(projectID, zone, instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to delete instance: %v", err)
	}

	waitVMOperation(service, projectID, zone, op.Name, "deleting")
	fmt.Printf("Instance %s deleted successfully!\n", instanceName)
	return nil
}

// func changeInstanceState(service *compute.Service, projectID, zone, instanceName, state string) error {
// 	ctx := context.Background()

// 	// Create a new Compute Service client
// 	service, err := compute.NewService(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// Check the current instance state and perform the desired state change
// 	switch state {
// 	case "START":
// 		opr, err := service.Instances.Start(projectID, zone, instanceName).Do()
// 		if err != nil {
// 			return err
// 		}
// 		return waitVMOperation(service, projectID, zone, opr.Name)

// 	case "STOP":
// 		opr, err := service.Instances.Stop(projectID, zone, instanceName).Do()
// 		if err != nil {
// 			return err
// 		}
// 		return waitVMOperation(service, projectID, zone, opr.Name)
// 	case "RESTART":
// 		opr, err := service.Instances.Reset(projectID, zone, instanceName).Do()
// 		if err != nil {
// 			return err
// 		}
// 		return waitVMOperation(service, projectID, zone, opr.Name)

// 	}

// 	return nil
// }

func SomeMultiplier(data int) {
	data *= data
	fmt.Println("Print the number square", data)
}
