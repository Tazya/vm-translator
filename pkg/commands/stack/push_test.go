package stack

import (
	"reflect"
	"testing"
)

func TestNewPush(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    *Push
		wantErr bool
	}{
		{
			name:    "Push from constant memory segment",
			command: "push constant 2",
			want: &Push{
				segment: "constant",
				index:   "2",
			},
		},
		{
			name:    "Incorrect arguments",
			command: "push unknown",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPush(tt.command)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPush_GetASMInstructions(t *testing.T) {
	type fields struct {
		segment string
		index   string
	}

	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Push from constant memory segment",
			fields: fields{
				segment: "constant",
				index:   "2",
			},
			want: []string{
				"// push constant 2",
				"@2",
				"D=A",
				"@SP",
				"A=M",
				"M=D",
				"@SP",
				"M=M+1",
			},
		},
		{
			name: "Push from incorrect memory segment",
			fields: fields{
				segment: "unknown",
				index:   "0",
			},
			want:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Push{
				segment: tt.fields.segment,
				index:   tt.fields.index,
			}
			got, err := p.GetASMInstructions()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetASMInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetASMInstructions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
