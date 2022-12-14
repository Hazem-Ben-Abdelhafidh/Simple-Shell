package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type History struct {
	command string
	time    string
}

// Create Directory
func createDir(names []string) {
	for _, name := range names {
		err := os.Mkdir(strings.TrimRight(name, "\r\n"), 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		fmt.Println("Directory ", name, " created!")
	}
}

// Create File
func createFile(name string) {
	_, err := os.Create(strings.TrimRight(name, "\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File ", name, " created!")
}

// Remove File/Directory
func remove(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File ", name, " Deleted!")
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var commandsHistory []History

forLoop:
	for {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("  ", strings.TrimSuffix(wd, "<nil>"))
		fmt.Printf("> ")

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
		case "mkdir":
			createDir(args[1:])
		case "cf":
			createFile(args[1])
		case "rm":
			remove(strings.TrimRight(args[1], "\r\n"))
		case "cd":
			os.Chdir(strings.TrimRight(args[1], "\r\n"))
		default:
			fmt.Println("unknown command!")

		}

	}

}
