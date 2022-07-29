package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type History struct {
	command string
	time    string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var commandsHistory []History
forLoop:
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		commandsHistory = append(commandsHistory, History{command: input, time: time.Now().Format("2006-01-02 15:04:05")})
		args := strings.Split(input, " ")
		switch strings.TrimRight(args[0], "\r\n") {

		case "exit":
			fmt.Println("Exiting...")
			break forLoop
		case "history":
			for i := 0; i < len(commandsHistory); i++ {
				fmt.Println(commandsHistory[i].time, " ", commandsHistory[i].command)
			}
		default:
			fmt.Println("unknown command!")

		}
	}

}
