package shell

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	DryRun         bool
}

func NewSession(in io.Reader, out, e io.Writer) Session {
	return Session{
		Stdin:  in,
		Stdout: out,
		Stderr: e,
	}
}

func (s *Session) Run() {
	fmt.Fprintf(s.Stdout, "> ")
	scanner := bufio.NewScanner(s.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cmd, e := CreateCommand(line)
		if e != nil {
			//fmt.Fprintf(s.Stderr, "error creating command %v\n", e)
			fmt.Fprintf(s.Stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s\n> ", line)
			continue
		}
		out, e := cmd.CombinedOutput()
		if e != nil {
			fmt.Fprintln(s.Stderr, "error:", e)
		}
		fmt.Fprintf(s.Stdout, "%s> ", out)
	}
	fmt.Fprintln(s.Stdout, "\nsee ya...")
}

func CreateCommand(cmd string) (*exec.Cmd, error) {
	args := strings.Fields(cmd)
	if len(args) < 1 {
		return nil, fmt.Errorf("cmd cant be empty%v", "")
	}

	return exec.Command(args[0], args[1:]...), nil

}
