package stack

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"strings"
	"testing"
)

func TestNewPop(t *testing.T) {
	tests := []struct {
		name      string
		command   string
		classname string
		want      *Pop
		wantErr   bool
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
			name:      "Pop to static memory segment",
			command:   "pop static 1",
			classname: "Mouse",
			want: &Pop{
				segment:   "static",
				index:     "1",
				classname: "Mouse",
			},
		},
		{
			name:    "Pop to wrong index in memory segment",
			command: "pop local 200000",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Pop to incorrect pointer memory segment",
			command: "pop pointer 2",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := strings.Fields(tt.command)
			segment := fields[1]
			index := fields[2]

			got, err := NewPop(segment, index, tt.classname)

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

	l := labels.NewLabels()
	tests := []struct {
		name      string
		fields    fields
		classname string
		want      []string
		wantErr   bool
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
				"@13",
				"M=D",
				"@SP", // @SP--
				"M=M-1",
				"A=M", // RAM[temp] = POP STACK
				"D=M",
				"@13",
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
				"@13",
				"M=D",
				"@SP", // @SP--
				"M=M-1",
				"A=M", // RAM[temp] = POP STACK
				"D=M",
				"@13",
				"A=M",
				"M=D",
			},
		},
		{
			name: "Pop to static memory segment",
			fields: fields{
				segment: "static",
				index:   "1",
			},
			classname: "Bird",
			want: []string{
				"// pop static 1",
				"@SP", // @SP--
				"M=M-1",
				"A=M",
				"D=M",
				"@Bird.1",
				"M=D",
			},
		},
		{
			name: "Pop to pointer 0 memory segment",
			fields: fields{
				segment: "pointer",
				index:   "0",
			},
			classname: "Bird",
			want: []string{
				"// pop pointer 0",
				"@SP", // @SP--
				"M=M-1",
				"A=M",
				"D=M",
				"@THIS",
				"M=D",
			},
		},
		{
			name: "Pop to pointer 1 memory segment",
			fields: fields{
				segment: "pointer",
				index:   "1",
			},
			classname: "Bird",
			want: []string{
				"// pop pointer 1",
				"@SP", // @SP--
				"M=M-1",
				"A=M",
				"D=M",
				"@THAT",
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
				segment:   tt.fields.segment,
				index:     tt.fields.index,
				classname: tt.classname,
			}
			got, err := p.GetASMInstructions(l)
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
