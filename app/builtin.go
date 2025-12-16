package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Loop() {
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
			executable := findExecutables(command[5:])
			if executable != "" {
				fmt.Printf("%v is %v", command[5:], executable)
				continue
			}
			fmt.Printf("%v: not found\n", command[5:])
			continue
		}

		fmt.Printf("%v: command not found\n", command)

	}
}
