package main

import (
	"log"
)

func main() {
	input := &StdinInput{}
	go func() {
		err := input.Listen()

		if err != nil {
			log.Fatal("SHUT IT DOWN!!!!!")
		}
	}()

	jsonParserMiddleware := &JSONParserMiddleware{}
	jsonParserMiddleware.Join(input.Subscribe())
	go jsonParserMiddleware.Listen()

	jsonMaskMiddleware := &JSONMaskMiddleware{}
	jsonMaskMiddleware.Join(jsonParserMiddleware.Subscribe())
	go jsonMaskMiddleware.Listen()

	jsonStringifierMiddleware := &JSONStringifierMiddleware{}
	jsonStringifierMiddleware.Join(jsonMaskMiddleware.Subscribe())
	go jsonStringifierMiddleware.Listen()

	output := &StdOutput{}

	output.Join(jsonStringifierMiddleware.Subscribe())
	output.Listen()
}
