package hello

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Command struct {
	outWriter io.Writer
	errWriter io.Writer
	// Flag and Args
	Name   string
	Format string
}

func New() *Command {
	return &Command{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}

func (c *Command) SetWriter(out, err io.Writer) {
	c.outWriter = out
	c.errWriter = err
}

func (c *Command) SetOutWriter(w io.Writer) {
	c.outWriter = w
}

func (c *Command) SetErrWriter(w io.Writer) {
	c.errWriter = w
}

func (c *Command) Command() *cobra.Command {
	command := &cobra.Command{
		Use:           "hello",
		Short:         "A brief description of command",
		Long:          `A longer description of command`,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: false,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			c.SetWriter(cmd.OutOrStdout(), cmd.ErrOrStderr())
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return c.Execute(ctx)
		},
	}

	command.Flags().StringVarP(&c.Name, "name", "n", "", "Name")
	command.Flags().StringVarP(&c.Format, "format", "f", "json", "Format")

	command.MarkFlagRequired("name")

	return command
}

func (c *Command) Execute(ctx context.Context) error {
	if (c.Format != "json") && (c.Format != "text") {
		return errors.New("invalid format")
	}

	if c.Format == "json" {
		return printHelloJSON(c.outWriter, c.Name)
	}
	if c.Format == "text" {
		return printHello(c.outWriter, c.Name)
	}

	return nil
}

func printHello(w io.Writer, name string) error {
	_, err := w.Write([]byte("Hello, " + name + "!\n"))
	return err
}

func printHelloJSON(w io.Writer, name string) error {
	_, err := w.Write([]byte(`{"message":"Hello, ` + name + `!"}`))
	return err
}
