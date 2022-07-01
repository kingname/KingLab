package runner

import (
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Shell   []string
	WorkDir string
}

func (c Command) Run() {
	cmd := exec.Command(c.Shell[0], c.Shell[1:]...)
	stdout := new(strings.Builder)
	stderr := new(strings.Builder)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Dir = c.WorkDir
	err := cmd.Run()
	if err != nil {
		fmt.Println("========")
		fmt.Println("xxxx", err)
	}
	fmt.Println("stdout：", stdout)
	fmt.Println("stderr: ", stderr)
}
