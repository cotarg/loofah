package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Input interface {
	Subscribe() chan string
	Listen() error
}

type Output interface {
	Join(chan string)
	Listen() error
}

type StringManglerMiddleware interface {
	Subscribe() chan string
	Join(chan string)
	Listen() error
}

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

type StdOutput struct {
	input chan string
}

func (x *StdOutput) Listen() error {
	// THESE BOTH DO THE SAME THING, WHICHEVER ONE IS COMMENTED IS NOT THE ONE I'M USING RIGHT NOW

	// for line := range x.input {
	// 	fmt.Println(line)
	// }

listenLoop:
	for {
		select {
		case line, ok := <-x.input:
			if !ok {
				break listenLoop
				// return nil
			}
			fmt.Println(line)
		}
	}
	return nil
}

func (x *StdOutput) Join(input chan string) {
	x.input = input
}

func (s *StdinInput) Subscribe() chan string {
	ch := make(chan string, 13)
	s.outputs = append(s.outputs, ch)
	return ch
}

func main() {
	input := &StdinInput{}
	go func() {
		err := input.Listen()

		if err != nil {
			log.Fatal("SHUT IT DOWN!!!!!")
		}
	}()

	output := &StdOutput{}

	output.Join(input.Subscribe())
	output.Listen()
}
