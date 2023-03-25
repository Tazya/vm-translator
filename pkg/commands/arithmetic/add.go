package arithmetic

type Add struct {
}

func (a *Add) GetASMInstructions() ([]string, error) {
	return []string{
		"// add",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=D+M",
		"@SP",
		"M=M+1",
	}, nil
}
