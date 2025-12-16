package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	builtinCommands := []string{"type", "exit", "echo"}

	for {
		var command string

		fmt.Print("$ ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			command = scanner.Text()
		}

		if command == "exit" {
			break
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}
		if strings.HasPrefix(command, "type ") {
			if slices.Contains(builtinCommands, command[5:]) {
				fmt.Printf("%v is a shell builtin\n", command[5:])
				continue
			}
			fmt.Printf("%v: not found\n", command[5:])
			continue
		}

		fmt.Printf("%v: command not found\n", command)

	}
}
