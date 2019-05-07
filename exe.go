package exe

import (
	"os"
	"os/exec"
	"strings"
)

type Exe struct {
	WorkingDir string
}

func New() *Exe {
	return &Exe{}
}

func (exe Exe) Run(command string) ([]string, error) {
	commandSegments := strings.Split(command, " ")
	cmd := exec.Command(commandSegments[0], commandSegments[1:]...)

	if exe.WorkingDir != "" {
		err := os.MkdirAll(exe.WorkingDir, 0700)
		if err != nil {
			return nil, err
		}
		cmd.Dir = exe.WorkingDir
	}

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(stdoutStderr), "\n")
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		return lines[0 : len(lines)-1], nil
	}
	return lines, nil
}
