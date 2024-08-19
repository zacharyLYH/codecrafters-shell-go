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
		"pwd": true,
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
		case strings.Contains(command, "pwd"):
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current working directory:", err)
				return
			}
		fmt.Println(pwd)
		default:
			splitCommand := strings.Split(command, " ")
			command := exec.Command(splitCommand[0], splitCommand[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", splitCommand[0])
			}
			// output, err := command.Output()
			// if err != nil {
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Printf("%s\n", string(output))
			// }
		}
	}
}
