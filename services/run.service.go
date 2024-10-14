package services

import (
	"fmt"
	"time"

	"watch-me/structs"

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
	data,_ := r.RawCommands.GetRunData()
	_, err := flags.Parse(data)
	if err != nil {
		fmt.Println("Error parsing args: ", err)
		return
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
