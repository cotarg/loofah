package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type JSONParserMiddleware struct {
	input   chan string
	outputs []chan JSONLogObject
}

type JSONLogObject struct {
	Date      string            `json:"date"`
	DateDay   string            `json:"date_day"`
	DateMonth string            `json:"date_month"`
	DateTime  string            `json:"date_time"`
	Hostname  string            `json:"hostname"`
	Message   string            `json:"message"`
	PID       string            `json:"pid"`
	Rig       map[string]string `json:"rig"`
	Syslog    map[string]string `json:"syslog"`
	Version   string            `json:"version"`
}

func (m *JSONParserMiddleware) Join(input chan string) {
	m.input = input
}

func (m *JSONParserMiddleware) Subscribe() chan JSONLogObject {
	ch := make(chan JSONLogObject, 13)
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

			modifiedLine, err := ingestLogline(line)

			if err != nil {
				log.Println("OH NOES! THIS NEEDS A BETTER MESSAGE! %v", err)
				continue
			}

			for _, output := range m.outputs {
				output <- modifiedLine
			}
		}
	}

	return nil
}

func ingestLogline(loglineString []byte) (JSONLogObject, error) {
	formattedLog := JSONLogObject{}
	err := json.Unmarshal(loglineString, &formattedLog)

	if err != nil {
		return formattedLog, fmt.Errorf("problem with decoding from JSON: %v", err)
	}

	return formattedLog, nil
}
