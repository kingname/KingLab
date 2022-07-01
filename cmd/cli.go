package main

import (
	"github.com/kingname/kinglab/internal/runner"
	"strings"
)

func main() {
	commander := runner.Command{
		Shell:   strings.Split("python3 test.py", " "),
		WorkDir: "/Users/kingname/Projects/KingLab/scripts",
	}

	commander.Run()
}
