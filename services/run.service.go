package services

import (
	"fmt"
	"time"

	"watch-me/structs"

	"watch-me/shared"

	"github.com/jessevdk/go-flags"
)

func generateRandomName() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("service-%d", timestamp)
}

func getRawCommandList() structs.CommandsData {
	commands := structs.CommandsRun{}
	return structs.CommandsData{
		Commands: &commands,
	}
}

func callback(r *structs.RunService) {
	data, _ := r.RawCommands.GetRunData()
	_, err := flags.Parse(data)
	if err != nil {
		fmt.Println("Error parsing args: ", err)
		return
	}
	fmt.Println("Running service")
	fmt.Println("Name: ", data.Name)
	fmt.Println("Dockerfile: ", data.DockerFile)
	fmt.Println("Dockerize: ", data.Dockerize)
	fmt.Println("CodeLang: ", data.CodeLang)
	if data.DockerFile != "" && data.Dockerize {
		panic("Cannot use both dockerfile and dockerize flags")
	}
	if data.Name == "" {
		data.Name = generateRandomName()
	}
	if data.CodeLang != "" {
		path := fmt.Sprintf("~/.config/watch-me/templates/Dockerfile.%s", data.CodeLang)
		fmt.Println("Path: ", path)
		switch data.CodeLang {
		case "node":
			shared.ExeCommand("cp", path, "./Dockerfile")
		}
	}
	if data.DockerFile != "" {
		shared.ExeCommand("docker", "build", "-t", data.Name, "-f", data.DockerFile, ".")
	} else {
		shared.ExeCommand("docker", "build", "-t", data.Name, ".")
	}
	shared.ExeCommand("docker", "run", "-d", "--name", data.Name+"-container", data.Name)

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
