package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the simple Shell!")
	var command string
	var commandsHistory []string
	for {
		fmt.Print(">")
		fmt.Scan(&command)
		commandsHistory = append(commandsHistory, command)
		args := strings.Split(command, " ")
		switch args[0] {
		case "cd":
			os.Chdir(args[1])
		case "history":
			for i := 0; i < len(commandsHistory); i++ {
				fmt.Println(commandsHistory[i])
			}

		default:
			fmt.Println("unknown command")
		}
	}
}
