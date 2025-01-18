package base

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Base struct {
	outWriter io.Writer
	errWriter io.Writer
}

func New() *Base {
	return &Base{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}

func (c *Base) PreRunE(cmd *cobra.Command, args []string) error {
	c.SetWriter(cmd.OutOrStdout(), cmd.ErrOrStderr())
	return nil
}

func (c *Base) SetWriter(out, err io.Writer) {
	c.outWriter = out
	c.errWriter = err
}

func (c *Base) SetOutWriter(w io.Writer) {
	c.outWriter = w
}

func (c *Base) SetErrWriter(w io.Writer) {
	c.errWriter = w
}

func (c *Base) Out() io.Writer {
	return c.outWriter
}

func (c *Base) Err() io.Writer {
	return c.errWriter
}
