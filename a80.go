package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type parser struct {
	// Global parser values
	lineNum int64  // Line number being parsed
	pass    int    // Which pass 1 or 2
	errors  bool   // Indicates an error
	output  []byte // Assembled code
	addr    int    // The address for labels
	// Parsed line values
	lab  string // Label from the current line
	op   string // Instruction mnemonic
	a1   string // First argument
	a2   string // Second argument
	comm string // Comment
}

func err(p parser, msg string) {
	fmt.Printf("a80: %d: %s", p.lineNum, msg)
	os.Exit(1)
}

func readFile(file string) ([]string, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getFilename(file string) string {
	return file[:len(file)-len(filepath.Ext(file))]
}

/**
 * Top-level assembly function.
 * Everything cascades downward from here.
 */
func assemble(p parser, lines []string, outfile string) {

}

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Printf("usage: a80 file.asm")
		return
	}

	lines, linesErr := readFile(args[1])
	if linesErr != nil {
		fmt.Printf("error reading file %s\n%s", args[1], linesErr)
		return
	}

	outfile := fmt.Sprintf("%s.com", getFilename(args[1]))
	p := parser{}

	fmt.Printf("Line count: %d / %s", len(lines), outfile)

	assemble(p, lines, outfile)
}

//
