package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kmio11/cobra-example/cmd"
	"github.com/kmio11/gilt"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd_Execute(t *testing.T) {
	golden := gilt.New[string, string](t.Name())

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name: "hello",
			args: []string{"hello", "--name", "world"},
		},
		{
			name:    "hello_invalid",
			args:    []string{"hello", "--name", "world", "--format", "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outWriter, errWriter bytes.Buffer

			root := cmd.NewRoot()
			root.SetOut(&outWriter)
			root.SetErr(&errWriter)
			root.SetArgs(tt.args)

			err := root.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("RootCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			golden.Assert(t, outWriter.String(), fmt.Sprintf("%s%s", tt.name, "_out"), func(t *testing.T, actual, expected string) {
				assert.Equal(t, expected, actual)
			})
			golden.Assert(t, errWriter.String(), fmt.Sprintf("%s%s", tt.name, "_err"), func(t *testing.T, actual, expected string) {
				assert.Equal(t, expected, actual)
			})
		})
	}
}
