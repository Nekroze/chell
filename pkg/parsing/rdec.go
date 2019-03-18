package parsing

import (
	"bufio"
	"io"
	"strings"
)

type Command struct {
	Executable string
	Arguments  []Argument
}

type Argument struct {
	Value string
}

type ASTnode interface {
	Parse(bufio.Scanner, string) (ASTnode, error)
}

type AST []ASTnode

type topLevel = Command

func Parse(s string) (ast AST, err error) {
	scanner := readerToScanner(strings.NewReader(s))

	for scanner.Scan() {
		node, err := topLevel{}.Parse(scanner, scanner.Text())
		if err != nil {
			return ast, err
		}
		ast = append(ast, node)
	}

	return ast, scanner.Err()
}

func (c Command) Parse(scanner bufio.Scanner, input string) (c ASTnode, err error) {
	err := TokIdentifier{}.Parse(input)
	if err != nil {
		return c, err
	}

	c.Executable = input

	return c, err
}

type TokIdentifier struct {
	Value string
}

func (ti TokIdentifier) Parse(input string) error {
	for _, r := range input {
		if !unicode.IsLetter(r) {
			return fmt.Errorf("'%s' should be alphabetical", input)
		}
	}
	return error
}

func readerToScanner(input io.Reader) bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Set the split function for the scanning operation.
	scanner.Split(splitter)
	return scanner
}

// Create a custom split function by wrapping the existing ScanWords function.
func splitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	return
}

type LexPart func(string) (string, error)

func lexNumber(input string) (found string, err error) {
	out := []rune{}
	for _, r := range input {
		if !unicode.IsNumber(r) {
			break
		}
		out = append(out, r)
	}

	if len(out) > 0 {
		err = fmt.Error("failed to parse number, %s contains illegal characters", input)
	}
	return string(out), err
}
