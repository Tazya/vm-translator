package stack

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewPop(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    *Pop
		wantErr bool
	}{
		{
			name:    "Pop to constant memory segment",
			command: "pop constant 2",
			want: &Pop{
				segment: "constant",
				index:   "2",
			},
		},
		{
			name:    "Pop to wrong index in memory segment",
			command: "pop local 200000",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := strings.Fields(tt.command)
			segment := fields[1]
			index := fields[2]

			got, err := NewPop(segment, index)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPop_GetASMInstructions(t *testing.T) {
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
			name: "Pop to local memory segment",
			fields: fields{
				segment: "local",
				index:   "2",
			},
			want: []string{
				"// pop local 2",
				"@LCL", // temp = local + 2
				"D=A",
				"@2",
				"D=D+A",
				"@5",
				"M=D",
				"@SP", // @SP--
				"M=M-1",
				"A=M", // RAM[temp] = POP STACK
				"D=M",
				"@5",
				"A=M",
				"M=D",
			},
		},
		{
			name: "Pop to argument memory segment",
			fields: fields{
				segment: "argument",
				index:   "4",
			},
			want: []string{
				"// pop argument 4",
				"@ARG", // temp = local + 2
				"D=A",
				"@4",
				"D=D+A",
				"@5",
				"M=D",
				"@SP", // @SP--
				"M=M-1",
				"A=M", // RAM[temp] = POP STACK
				"D=M",
				"@5",
				"A=M",
				"M=D",
			},
		},
		{
			name: "Pop to incorrect memory segment",
			fields: fields{
				segment: "constant",
				index:   "0",
			},
			want:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pop{
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
