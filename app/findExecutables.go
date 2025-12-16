package main

import (
	"fmt"
	"os"
	"strings"
)

func findExecutables(command string) string {
	path := os.Getenv("PATH")
	if path == "" {
		return ""
	}
	paths := strings.Split(path, ":")

	for _, dir := range paths {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			if !file.IsDir() && file.Name() == command {
				info, err := file.Info()
				if err != nil {
					continue
				}
				isExecutable := info.Mode().Perm()&0o100 != 0
				if isExecutable {
					return fmt.Sprintf("%v/%v", dir, file.Name())
				}

			}
		}
	}

	return ""
}
