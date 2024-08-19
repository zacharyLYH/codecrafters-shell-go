package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func checkIfIsExecutable(name string, paths []string) string {
	for _, p := range paths{
		exec := p + "/" + name
		fmt.Println(exec)
		if _, err := os.Stat(exec); err == nil {
			return exec
		}
	}
	return ""
}

func main() {
	builtInCommands := map[string]interface{}{
		"echo": true,
		"exit": true,
		"type": true,
	}
	paths := strings.Split(os.Getenv("PATH"), ":")
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		switch {
		case command == "exit 0":
			return
		case strings.Contains(command, "echo "):
			fmt.Println(strings.TrimPrefix(command, "echo "))
		case strings.Contains(command, "type "):
			check := strings.TrimPrefix(command, "type ")
			if _,exists := builtInCommands[check]; exists{
				fmt.Printf("%s is a shell builtin\n",check)
			}else{
				executablePath := checkIfIsExecutable(check, paths)
				if executablePath == ""{
					fmt.Printf("%s: not found\n", check)
				} else {
					fmt.Fprintf(os.Stdout, "%v is %v\n", check, executablePath)
				}
			}
		default:
			progName := os.Args[0]
			executablePath := checkIfIsExecutable(progName, paths)
			if executablePath == ""{
				fmt.Printf("%s: not found\n", progName)
			} else {
				args := os.Args[1:]
				cmd := exec.Command(progName, args...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			}
		}
	}
}
