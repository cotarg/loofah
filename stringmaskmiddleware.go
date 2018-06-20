package main

import (
	"fmt"
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
		case line, ok := <-x.input:
			if !ok {
				break listenLoop
			}
			fmt.Println(strings.ToUpper(line))
		}
	}
	return nil
}

func (m *StringMaskMiddleware) Subscribe() chan string {
	ch := make(chan string, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}
