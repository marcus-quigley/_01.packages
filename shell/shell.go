package shell

import (
	"fmt"
	"os/exec"
	"strings"
)

func CreateCommand(cmd string) (*exec.Cmd, error) {
	args := strings.Fields(cmd)
	if len(args) < 1 {
		return nil, fmt.Errorf("cmd cant be empty%v", "")
	}

	return exec.Command(args[0], args[1:]...), nil

}
