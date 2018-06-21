package main

import (
	"strings"
)

type StringMaskMiddleware struct {
	input   chan string
	outputs []chan string
}

func (m *StringMaskMiddleware) Join(input chan string) {
	m.input = input
}

func (m *StringMaskMiddleware) Listen() error {
listenLoop:
	for {
		select {
		case line, ok := <-m.input:
			if !ok {
				break listenLoop
			}

			for _, output := range m.outputs {
				output <- strings.ToUpper(line)
			}
		}
	}

	// close the channel because NO MOAR STUFF
	for _, output := range m.outputs {
		close(output)
	}
	return nil
}

func (m *StringMaskMiddleware) Subscribe() chan string {
	ch := make(chan string, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}
