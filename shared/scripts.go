package shared

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)
func ExeCommand(args ...string) (string, error) {
    // Resolve any "~" in the args to the home directory
    for i, arg := range args {
        if strings.HasPrefix(arg, "~") {
            homeDir, err := os.UserHomeDir()
            if err != nil {
                fmt.Println("Error resolving home directory:", err)
                return "", err
            }
            args[i] = strings.Replace(arg, "~", homeDir, 1)
        }
    }

    cmd := exec.Command(args[0], args[1:]...)

    var stdoutBuf, stderrBuf bytes.Buffer

    stdoutMulti := io.MultiWriter(os.Stdout, &stdoutBuf)
    stderrMulti := io.MultiWriter(os.Stderr, &stderrBuf)

    cmd.Stdout = stdoutMulti
    cmd.Stderr = stderrMulti

    err := cmd.Run()

    logs := stdoutBuf.String() + stderrBuf.String()

    if err != nil {
        fmt.Println("Error running command:", err)
        return logs, err
    }

    return logs, nil
}
func GetCallingPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return cwd
}
