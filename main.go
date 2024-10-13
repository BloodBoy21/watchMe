package main

import (
	"flag"
	"fmt"

	"watch-me/services"
)


func main() {
    cli := services.CLI{}
    cli.Init()
    flag.Parse()
    fmt.Printf("Flags: %d\n", flag.NFlag())
     args := flag.Args()

    // Handle the subcommand (if any)
    if len(args) > 0 && args[0] == "run" {
        fmt.Println("Subcommand: run")
    }

    cli.Run()
}