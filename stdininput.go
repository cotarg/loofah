package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type StdinInput struct {
	outputs []chan string
}

func (s *StdinInput) Listen() error {
	// read in stdin
	statsInfo, err := os.Stdin.Stat()
	reader := bufio.NewReader(os.Stdin)

	// if there's an err at this point, you failed reading std in
	if err != nil {
		return fmt.Errorf("omg something's wrong with reading stdin: %v", err)
	}

	// if the number of inputs is wrong, boot out
	if statsInfo.Mode()&os.ModeNamedPipe == 0 {
		return errors.New("either too much or not enough input womp womp!")
	}

	for {
		// read line by line
		line, err := reader.ReadString('\n')

		// if end of file, just break, else shout
		if err != nil {
			if err == io.EOF {
				break
			}

			// if err != EOF, tell me what's wrong
			return fmt.Errorf("omg something's wrong with reading stdin: %v", err)
		}

		// reader.ReadString() retains delimiter, so this strips it for output
		line = strings.TrimSuffix(line, "\n")

		// fmt.Println(line)
		for _, output := range s.outputs {
			output <- line
		}
	}

	// close the channel because NO MOAR STUFF
	for _, output := range s.outputs {
		close(output)
	}

	return nil
}

func (s *StdinInput) Subscribe() chan string {
	ch := make(chan string, 13)
	s.outputs = append(s.outputs, ch)
	return ch
}
