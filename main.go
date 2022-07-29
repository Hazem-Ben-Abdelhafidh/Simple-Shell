package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type commandName string

const (
	// change Directory
	ChangeDirectory commandName = "cd"
	CreateDirectory commandName = "cr"
	CommandsHistory commandName = "history"
	CreateFile      commandName = "cf"
	ReadFile        commandName = "rf"
	Exit            commandName = "close"
)

type Command struct {
	Name commandName
	Args []string
}

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	cmd := exec.Command(input)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	fmt.Println("Welcome to the simple shell!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
