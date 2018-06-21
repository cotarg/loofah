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
	defer func() {
		// close the channel because NO MOAR STUFF
		for _, output := range m.outputs {
			close(output)
		}
	}()

	for {
		select {
		case line, ok := <-m.input:
			if !ok {
				return nil
			}

			for _, output := range m.outputs {
				output <- emojifier(line)
			}
		}
	}

	return nil
}

func (m *StringMaskMiddleware) Subscribe() chan string {
	ch := make(chan string, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}

// emojifier converts a's into a cat heart eye emoji
func emojifier(line string) string {
	find := 'a'
	replace := '😻'
	return strings.Map(func(r rune) rune {
		if r == find {
			return replace
		}
		return r
	}, line)
}
