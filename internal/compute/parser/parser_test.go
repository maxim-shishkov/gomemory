package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute_Parse(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    Query
		wantErr bool
	}{
		{
			name: "set",
			text: "SET key value",
			want: Query{
				command: CommandSET,
				args:    []string{"key", "value"},
			},
			wantErr: false,
		},
		{
			name: "get",
			text: "GET key",
			want: Query{
				command: CommandGET,
				args:    []string{"key"},
			},
			wantErr: false,
		},
		{
			name: "set",
			text: "DEL key",
			want: Query{
				command: CommandDEL,
				args:    []string{"key"},
			},
			wantErr: false,
		},
		{
			name:    "unknown",
			text:    "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compute := &Compute{
				logger: nil,
			}
			got, err := compute.Parse(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
