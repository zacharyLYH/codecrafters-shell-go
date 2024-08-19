package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		switch {
		case command == "exit 0":
			return
		case strings.Contains(command, "echo "):
			fmt.Println(strings.TrimPrefix(command, "echo "))
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
