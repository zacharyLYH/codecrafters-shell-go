package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
				isExecutable := false
				for _, p := range paths{
					exec := p + "/" + check
					if _, err := os.Stat(exec); err == nil {
						fmt.Fprintf(os.Stdout, "%v is %v\n", check, exec)
						isExecutable = true
						break
					}
				}
				if !isExecutable{
					fmt.Printf("%s: not found\n", check)
				}
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
