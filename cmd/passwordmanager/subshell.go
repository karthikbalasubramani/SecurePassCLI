package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func StartShell() {
	// Create a new scanner to read user input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	// Process user commands until the shell is exited
	fmt.Println("The available commands are:")
	HelpHandler(nil, nil)
	// SigninHandler(nil, nil)
	for {
		fmt.Print(">> ")
		// Read the user input
		scanner.Scan()
		input := scanner.Text()
		// Split the input into command and arguments
		parts := strings.Split(input, " ")
		if len(parts) == 0 {
			continue
		}
		// Extract the command and arguments
		command := parts[0]
		switch command {
		case "add":
			AddHandler(nil, nil)
		case "help":
			HelpHandler(nil, nil)
		case "signup":
			SignupHandler(nil, nil)
		case "get":
			GetHandler(nil, nil)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}