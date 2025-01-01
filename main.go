package main

import (
	"context"
	"os"

	"github.com/kmio11/cobra-example/cmd"
)

func main() {
	ctx := context.Background()

	rootCmd := cmd.NewRoot()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
