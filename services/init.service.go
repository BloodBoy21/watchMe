package services

import (
	"watch-me/shared"
	"watch-me/structs"
)



func initCallback(r *structs.RunService) {
	/* 	data, _ := r.RawCommands.GetInitData()
		_, err := flags.Parse(data)
		if err != nil {
			fmt.Println("Error parsing args: ", err)
			return
		} */
	shared.	ExeCommand("mkdir", "-p", "~/.config/watch-me/templates")
	shared.	ExeCommand("cp","-r" ,"./templates", "~/.config/watch-me")

}

func NewInitService() *structs.RunService {
	_model := structs.CommandModel{}
	commands := structs.CommandsRun{}
	rawCommands := structs.CommandsData{
		Commands: &commands,
	}
	service := &structs.RunService{
		Model:        _model,
		EntryCommand: "init",
		Command:      "init",
		RawCommands:  rawCommands,
		Callback:     initCallback,
	}
	return service
}
