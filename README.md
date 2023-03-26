# Virtual Machine Translator for Nand to Tetris course

[![Build and test](https://github.com/Tazya/vm-translator/actions/workflows/tests.yml/badge.svg)](https://github.com/Tazya/vm-translator/actions/workflows/tests.yml)

VMTranslator converts virtual machine code to HACK-assembly language

Usage:
```
 go run main.go -fileIn="./examples/Pop.vm" -dirOut="./output/"
```

This command converts the vm-code:
```
push constant 8
pop local 2
```

To a bunch of asm instructions for the HACK computer:
```
// push constant 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop local 2
@LCL
D=A
@2
D=D+A
@5
M=D
@SP
M=M-1
A=M
D=M
@5
A=M
M=D
```