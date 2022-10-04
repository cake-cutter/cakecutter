package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func runCommand(command string, stdout, stderr *bytes.Buffer) error {
	cmands, err := shellwords.Parse(command)
	if err != nil {
		return err
	}
	if len(cmands) == 0 {
		return fmt.Errorf("empty command")
	}
	cmd := exec.Command(cmands[0], cmands[1:]...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("stderr message: %s, status: %w", stderr.String(), err)
	}
	return nil
}
