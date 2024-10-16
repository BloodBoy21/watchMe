package structs

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type CommandsData struct {
	Commands any
}

type FlagValue struct {
	Str  *string
	Bool *bool
}

type RunService struct {
	Model        CommandModel
	EntryCommand string
	Command      string
	RawCommands  CommandsData
	Callback     func(*RunService)
	UseTea       bool
	TeaCallback  func(tea.Msg) (tea.Model, tea.Cmd)
}

func (cd *CommandsData) GetRunData() (*CommandsRun, error) {
	if runData, ok := cd.Commands.(*CommandsRun); ok {
		return runData, nil
	}
	return nil, fmt.Errorf("Commands is not of type *CommandsRun")
}

func (cd *CommandsData) GetInitData() (*CommandsInit, error) {
	if initData, ok := cd.Commands.(*CommandsInit); ok {
		return initData, nil
	}
	return nil, fmt.Errorf("Commands is not of type *CommandsInit")
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

