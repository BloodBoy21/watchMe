package shared

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExeCommand(args ...string) {

    // Expand '~' to the home directory if used in the path
    for i, arg := range args {
        if strings.HasPrefix(arg, "~") {
            homeDir, err := os.UserHomeDir()
            if err != nil {
                fmt.Println("Error resolving home directory:", err)
                return
            }
            args[i] = strings.Replace(arg, "~", homeDir, 1)
        }
    }

    // Prepare the command
    cmd := exec.Command(args[0], args[1:]...)

    // Set stdout and stderr to os.Stdout and os.Stderr to stream output in real time
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Run the command and wait for it to finish
    err := cmd.Run()

    if err != nil {
        fmt.Printf("Error running command: %v\n", err)
    }
}
