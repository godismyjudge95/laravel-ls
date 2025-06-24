package runtime

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"strings"
)

// PHPProcess represents a PHP execution process with configurable arguments.
type PHPProcess struct {
	Args []string
}

// NewPHPProcess creates a new PHPProcess instance with the given arguments.
func NewPHPProcess(args ...string) *PHPProcess {
	return &PHPProcess{
		Args: args,
	}
}

// Exec executes the PHP code in the specified working directory.
// It returns an io.Reader with the output or an error if execution fails.
func (proc PHPProcess) Exec(workingDir string, code []byte) (io.Reader, error) {
	outBuf := &bytes.Buffer{}
	errBuf := &strings.Builder{}

	// Prepare the command with the PHP binary and code as arguments.
	// proc.Args[0] is the PHP binary (e.g., "php"), and the rest are additional arguments.
	cmd := exec.Command(proc.Args[0], append(proc.Args[1:], string(code))...)
	cmd.Dir = workingDir
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	if cmdErr := cmd.Run(); cmdErr != nil {
		err := errors.New(errBuf.String())
		return nil, errors.Join(err, cmdErr)
	}
	return outBuf, nil
}
