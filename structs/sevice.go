package structs

import (
	"flag"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type FlagValue struct {
	Str  *string
	Bool *bool
}

type RunService struct {
	Model        CommandModel
	CommandsList map[string]*interface{}
	Commands     map[string]Command
	EntryCommand string
	Command      string
	RawCommands  []Command
	Callback     func(*RunService)
	UseTea       bool
	TeaCallback  func(tea.Msg) (tea.Model, tea.Cmd)
}

func (f *FlagValue) Called(command Command) bool {
	if command.Type == "bool" {
		return *f.Bool
	}
	return f.Str != &command.DefaultValue
}

func (r *RunService) Run() {
	r.Callback(r)
	if r.UseTea {
		p := tea.NewProgram(r.Model)
		if r.TeaCallback != nil {
			r.Model.UpdateCallback = r.TeaCallback
		}
		if _, err := p.Run(); err != nil {
			panic(err)
		}
	}
}
func (r *RunService) InitCommands() {
	// Initialize the maps if they are nil
	if r.Commands == nil {
		r.Commands = make(map[string]Command)
	}
	if r.CommandsList == nil {
		r.CommandsList = make(map[string]*interface{}) // Ensure CommandsList is initialized
	}

	for _, command := range r.RawCommands {
		r.Commands[command.Name] = command
		switch command.Type {
		case "bool":
			val := flag.Bool(command.Flag, command.ParseDefault().(bool), command.Description)
			var valInterface interface{} = val
			r.CommandsList[command.Name] = &valInterface
		case "string":
			val := flag.String(command.Flag, command.ParseDefault().(string), command.Description)
			var valInterface interface{} = val
			r.CommandsList[command.Name] = &valInterface
		default:
			fmt.Println("Invalid command type")
		}
	}
}
