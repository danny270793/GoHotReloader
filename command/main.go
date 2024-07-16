package command

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	command string
	path    string

	cmd    *exec.Cmd
	stdout io.ReadCloser
	cancel context.CancelFunc
}

func New(path string, command string) Command {
	return Command{command: command, path: path}
}

func (c *Command) Cancel() {
	c.cancel()
}

func (c *Command) Run() error {
	parts := strings.Fields(c.command)
	cmdName := parts[0]
	cmdArgs := parts[1:]

	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	cmd := exec.CommandContext(ctx, cmdName, cmdArgs...)
	cmd.Dir = c.path
	c.cmd = cmd

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("StdoutPipe not created because: %v", err)
	}
	c.stdout = stdout

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("command not started because: %v", err)
	}

	buffer := make([]byte, 4096)
	for {
		n, err := stdout.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("stdout not read because: %v", err)
		}

		fmt.Printf("%s", buffer[:n])
	}

	return nil
}
