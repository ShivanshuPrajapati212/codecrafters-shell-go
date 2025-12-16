package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
				fmt.Printf("%v is %v\n", command[5:], executable)
				continue
			}
			fmt.Printf("%v: not found\n", command[5:])
			continue
		}

		args := strings.Split(command, " ")
		var remainingArgs []string
		if len(args) > 2 {
			remainingArgs = args[1:]
		}

		executable := findExecutables(args[0])
		if executable != "" {
			cmd := exec.Command(executable, remainingArgs...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Print(err, "\n")
				continue
			}
			fmt.Print(string(output))
			cmd.Run()
			continue
		}

		fmt.Printf("%v: command not found\n", command)

	}
}
