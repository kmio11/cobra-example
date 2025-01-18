package hello

import (
	"context"
	"errors"
	"io"

	"github.com/kmio11/cobra-example/base"
	"github.com/spf13/cobra"
)

type Command struct {
	*base.Base
	// Flag and Args
	Name   string
	Format string
}

func New() *Command {
	return &Command{
		Base: base.New(),
	}
}

func (c *Command) Command() *cobra.Command {
	command := &cobra.Command{
		Use:           "hello",
		Short:         "A brief description of command",
		Long:          `A longer description of command`,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: false,
		PreRunE:       c.PreRunE,
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
		return printHelloJSON(c.Out(), c.Name)
	}
	if c.Format == "text" {
		return printHello(c.Out(), c.Name)
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
