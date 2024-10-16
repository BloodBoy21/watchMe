package main

import (
	"watch-me/services"
)


func main() {
    cli := services.CLI{}
    cli.Init()
    cli.Run()
}