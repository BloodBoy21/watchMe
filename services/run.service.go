package services

import (
	"fmt"
	"time"

	"watch-me/structs"
)

func generateRandomName() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("service-%d", timestamp)
}

func getRawCommandList() []structs.Command {
	return []structs.Command{
		{
			Name:         "UseDockerFile",
			Description:  "Use a Dockerfile",
			Type:         "bool",
			DefaultValue: "false",
			Flag:         "d",
			Weight:       1,
		},
		{
			Name:         "DockerFile",
			Description:  "Dockerfile path",
			Type:         "string",
			DefaultValue: "",
			Flag:         "f",
			Weight:       2,
		},
		{
			Name:         "Type",
			Description:  "Service type (node, python, go, ...)",
			Type:         "string",
			DefaultValue: "",
			Flag:         "t",
			Weight:       1,
		},
		{
			Name:         "Name",
			Type:         "string",
			Description:  "Service name",
			DefaultValue: generateRandomName(),
			Flag:         "n",
			Weight:       3,
		},
	}
}

func groupByWeight(commands []structs.Command) map[int][]structs.Command {
	grouped := make(map[int][]structs.Command)
	for _, command := range commands {
		grouped[command.Weight] = append(grouped[command.Weight], command)
	}
	return grouped
}

func callback(r *structs.RunService) {
	weightGroupedCommands := groupByWeight(r.RawCommands)
	for _, commands := range weightGroupedCommands {
		fmt.Printf("\nWeight: %d\n", commands[0].Weight)
		for _, command := range commands {
			flagCommand := r.CommandsList[command.Name]
			fmt.Printf("Name: %s\n", command.Name)
			fmt.Printf("Description: %s\n", command.Description)
			fmt.Printf("Type: %s\n", command.Type)
			fmt.Printf("Flag: %s\n", command.Flag)
			fmt.Printf("Weight: %d\n", command.Weight)
			// Handle bool and string types with pointer type assertion
			if command.Type == "bool" {
				val := (*flagCommand).(*bool) // Type assert to *bool, then dereference
				fmt.Printf("Value: %t\n", *val)
			} else if command.Type == "string" {
				val := (*flagCommand).(*string) // Type assert to *string, then dereference
				fmt.Printf("Value: %s\n", *val)
			}
		fmt.Println("------------------------------------------------")
		}
	}
}

func NewRunService() *structs.RunService {
	_model := structs.CommandModel{}
	service := &structs.RunService{
		Model:        _model,
		EntryCommand: "run",
		Command:      "run",
		RawCommands:  getRawCommandList(),
		Callback:     callback,
	}
	return service
}
