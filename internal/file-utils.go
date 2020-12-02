package internal

import (
	"bytes"
	"io/ioutil"
)

// reads a file and returns an array of strings where each element is a line in the file
func ReadFile(filename string) ([]string, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rawInput = bytes.TrimRight(rawInput, "\n")
	linesAsBytes := bytes.Split(rawInput, []byte("\n"))

	lines := make([]string, 0)
	for _, l := range linesAsBytes {
		lines = append(lines, string(l))
	}

	return lines, nil
}