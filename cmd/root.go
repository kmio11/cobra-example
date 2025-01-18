package cmd

import (
	"context"

	"github.com/kmio11/cobra-example/cmd/hello"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	*cobra.Command
}

func NewRoot() *cobra.Command {
	rootCmd := &RootCmd{
		Command: &cobra.Command{
			Use:   "cobra-example",
			Short: "A brief description of your application",
			Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		},
	}

	rootCmd.AddCommand(hello.New().Command())

	return rootCmd.Command
}

func (cmd *RootCmd) Execute() error {
	return cmd.Command.Execute()
}

func (cmd *RootCmd) ExecuteContext(ctx context.Context) error {
	return cmd.Command.ExecuteContext(ctx)
}
