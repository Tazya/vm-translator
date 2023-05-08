package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tazya/vm-translator/pkg/labels"
	"github.com/tazya/vm-translator/pkg/parser"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var inputFilepath string
	var outputDirectory string

	flag.StringVar(&inputFilepath, "fileIn", "", "a string var")
	flag.StringVar(&outputDirectory, "dirOut", "", "a string var")
	flag.Parse()

	if inputFilepath == "" || outputDirectory == "" {
		fmt.Println("Usage: vm-translator -fileIn=\"path/to/input.vm\" dirOut=\"path/to/output\"")
		return
	}

	inputFile, err := os.Open(inputFilepath)

	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	fileNameWithoutExt, _, _ := strings.Cut(filepath.Base(inputFilepath), ".")
	outputPath := outputDirectory + "/" + fileNameWithoutExt + ".asm"

	f, err := os.Create(outputPath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	l := labels.NewLabels()
	inputFileScanner := bufio.NewScanner(inputFile)
	inputFileScanner.Split(bufio.ScanLines)

	codeLine := 0

	for inputFileScanner.Scan() {
		codeLine++

		command, err := parser.ParseLine(inputFileScanner.Text(), fileNameWithoutExt)

		if err != nil {
			fmt.Println(fmt.Sprintf("Parsing error. line:%d", codeLine), err.Error())

			return
		}

		if command == nil {
			continue
		}

		instructions, err := command.GetASMInstructions(l)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error. line:%d", codeLine), err.Error())

			return
		}

		for _, instruction := range instructions {
			f.WriteString(instruction + "\n")
		}
	}
}
