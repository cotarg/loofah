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

	middleware := &StringMaskMiddleware{}
	middleware.Join(input.Subscribe())
	go middleware.Listen()

	output := &StdOutput{}

	output.Join(middleware.Subscribe())
	output.Listen()
}
