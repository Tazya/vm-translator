package stack

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewPush(t *testing.T) {
	tests := []struct {
		name      string
		command   string
		classname string
		want      *Push
		wantErr   bool
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
			name:      "Push from static memory segment",
			command:   "push static 1",
			classname: "cat",
			want: &Push{
				segment:   "static",
				index:     "1",
				classname: "cat",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := strings.Fields(tt.command)
			segment := fields[1]
			index := fields[2]

			got, err := NewPush(segment, index, tt.classname)

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
		name      string
		fields    fields
		classname string
		want      []string
		wantErr   bool
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
			name: "Push from local memory segment",
			fields: fields{
				segment: "local",
				index:   "0",
			},
			classname: "cat",
			want: []string{
				"// push local 0",
				"@LCL",
				"D=M",
				"@0",
				"A=D+A",
				"D=M",
				"@SP",
				"A=M",
				"M=D",
				"@SP",
				"M=M+1",
			},
		},
		{
			name: "Push from static memory segment",
			fields: fields{
				segment: "static",
				index:   "3",
			},
			classname: "cat",
			want: []string{
				"// push static 3",
				"@cat.3",
				"D=M",
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
				segment:   tt.fields.segment,
				index:     tt.fields.index,
				classname: tt.classname,
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
