package services

import (
	"fmt"
	"os"

	"watch-me/structs"
)

type CLI struct {
	modules map[string]*structs.RunService
}

func (c *CLI) init() {
	c.modules = map[string]*structs.RunService{
		"run": NewRunService(),
	}
	for _, module := range c.modules {
		module.InitCommands()
	}
}

func (c *CLI) verifyEntryCommand(command string) (string,error) {
	fmt.Printf("command: %s\n", command)
	for key := range c.modules {
		if key == command {
			return key, nil
		}
	}
	return "", fmt.Errorf("Invalid entry command: %s", command)
}

func (c *CLI) Init() *CLI{
	c.init()
	return c
}

func (c *CLI) GetEntryCommand() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}
	entryCommand, err := c.verifyEntryCommand(args[1])
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return entryCommand
}

func (c *CLI) Run() {
	fmt.Println("Running CLI")
	entryCommand := c.GetEntryCommand()
	if entryCommand == "" {
		fmt.Println("Invalid entry command")
		return
	}
	module := c.modules[entryCommand]
	if module == nil {
		fmt.Println("Invalid module")
		return
	}
	module.Run()
}