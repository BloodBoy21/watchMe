package services

import (
	"fmt"
	"time"

	"github.com/jessevdk/go-flags"

	"watch-me/shared"
	"watch-me/structs"
)

func generateRandomName() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("service-%d", timestamp)
}

func runCallback(r *structs.RunService) {
	data, _ := r.RawCommands.GetRunData()
	_, err := flags.Parse(data)
	if err != nil {
		fmt.Println("Error parsing args: ", err)
		return
	}
	fmt.Println("Name: ", data.Name)
	fmt.Println("Dockerfile: ", data.DockerFile)
	fmt.Println("Dockerize: ", data.Dockerize)
	fmt.Println("CodeLang: ", data.CodeLang)
	fmt.Println()
	callingPath := shared.GetCallingPath()
	if data.DockerFile != "" && data.Dockerize {
		panic("Cannot use both dockerfile and dockerize flags")
	}
	if data.Name == "" {
		data.Name = generateRandomName()
	}
	if data.CodeLang != "" {
		dockerfilePath := fmt.Sprintf("~/.config/watch-me/templates/Dockerfile.%s", data.CodeLang)
		dockerIgnoresPath := fmt.Sprintf("~/.config/watch-me/templates/ignores/dockerignore.%s", data.CodeLang)
		switch data.CodeLang {
		case "node":
			shared.ExeCommand("cp", dockerfilePath, "./Dockerfile")
			shared.ExeCommand("cp", dockerIgnoresPath, "./.dockerignore")
		}
		data.DockerFile = callingPath + "/Dockerfile"
	}
	_, err = shared.ExeCommand("docker", "build", "-t", data.Name, "-f", data.DockerFile, callingPath, "--no-cache")
	if err != nil {
		panic(err)
	}
	dockerId, err := shared.ExeCommand("docker", "run", "-d", "--name", data.Name+"-container", data.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Docker container id: ", dockerId)
}

func NewRunService() *structs.RunService {
	_model := structs.CommandModel{}
	commands := structs.CommandsRun{}
	rawCommands := structs.CommandsData{
		Commands: &commands,
	}
	service := &structs.RunService{
		Model:        _model,
		EntryCommand: "run",
		Command:      "run",
		RawCommands:  rawCommands,
		Callback:     runCallback,
	}
	return service
}
