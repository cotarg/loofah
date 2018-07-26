package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type JSONStringifierMiddleware struct {
	input   chan JSONLogObject
	outputs []chan string
}

func (m *JSONStringifierMiddleware) Join(input chan JSONLogObject) {
	m.input = input
}

func (m *JSONStringifierMiddleware) Subscribe() chan string {
	ch := make(chan string, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}

func (m *JSONStringifierMiddleware) Listen() error {
	defer func() {
		// close the channel because NO MOAR STUFF
		for _, output := range m.outputs {
			close(output)
		}
	}()

	for {
		select {
		case jsonObj, ok := <-m.input:
			if !ok {
				return nil
			}

			stringifiedJSON, err := ingestJSONObj(jsonObj)

			if err != nil {
				log.Println("OH NOES! THIS NEEDS A BETTER MESSAGE! %v", err)
				continue
			}

			for _, output := range m.outputs {
				output <- stringifiedJSON
			}
		}
	}

	return nil
}

func ingestJSONObj(jsonLogObjToStringify JSONLogObject) (string, error) {
	stringifiedJSON, err := json.Marshal(&jsonLogObjToStringify)

	if err != nil {
		return "", fmt.Errorf("A FRAYED KNOT! Problem with making a string: %v", err)
	}

	return string(stringifiedJSON), nil
}
