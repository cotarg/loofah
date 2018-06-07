package main

import (
	"bufio"
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

type StdinInput struct {
	outputs []chan string
}

func (s *StdinInput) Listen() error {
	statsInfo, err := os.Stdin.Stat()
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Fatal("OH NOES THIS IS SO BROKEN!!")
	}

	if statsInfo.Mode()&os.ModeNamedPipe == 0 {
		log.Fatal("TOO MANY THINGS OR NOT ENOUGH THINGS BUT THINGS ARE WRONG!")
	}

	for {
		line, err := reader.ReadString('\n')

		// if end of file, just break, else shout
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal("OH NOES THIS IS SO BROKEN!!")

		}

		// reader.ReadString() retains delimiter, so this strips it for output
		line = strings.TrimSuffix(line, "\n")

		fmt.Println(line)
	}
}

func (s *StdinInput) Subscribe() chan string {
	ch := make(chan string, 13)
	s.outputs = append(s.outputs, ch)
	return ch
}

func main() {

}
