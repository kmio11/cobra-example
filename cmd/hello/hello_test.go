package hello_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/kmio11/cobra-example/cmd/hello"
	"github.com/kmio11/gilt"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Execute(t *testing.T) {

	golden := gilt.New[string, string](t.Name())

	tests := []struct {
		name    string
		format  string
		wantErr bool
	}{
		{
			name:   "json",
			format: "json",
		},
		{
			name:   "text",
			format: "text",
		},
		{
			name:    "invalid",
			format:  "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var outWriter, errWriter bytes.Buffer

			cmd := hello.New()
			cmd.SetWriter(&outWriter, &errWriter)
			cmd.Name = tt.name
			cmd.Format = tt.format

			err := cmd.Execute(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Command.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}

			actual := outWriter.String()
			golden.Assert(t, actual, tt.name, func(t *testing.T, actual, expected string) {
				assert.Equal(t, expected, actual)
			})
		})
	}
}
