package main

import (
	"encoding/json"
	"fmt"
)

type JSONParserMiddleware struct {
	input   chan string
	outputs []chan string
}

type JSONLogObject struct {
	date       string
	date_day   string
	date_month string
	date_time  string
	hostname   string
	message    string
	pid        string
	rig        map[string]string
	syslog     map[string]string
	version    string
}

func (m *JSONParserMiddleware) Join(input chan string) {
	m.input = input
}

func (m *StringMaskMiddleware) Subscribe() chan string {
	ch := make(chan string, 13)
	m.outputs = append(m.outputs, ch)
	return ch
}

func (m *JSONParserMiddleware) Listen() error {
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
				line = []byte(line)
				output <- logScrub(line)
			}
		}
	}

	return nil
}

func logScrub(formattedLog JSONLogObject) {

}

func ingestLogline(loglineString []byte) JSONLogObject {
	var formattedLog JSONLogObject
	err = json.Unmarshal(loglineString, &formattedLog)

	if err != nil {
		return fmt.Errorf("problem with decoding from JSON: %v", err)
	}

	return formattedLog
}
