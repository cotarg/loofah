package main

import "fmt"

type StdOutput struct {
	input chan string
}

func (x *StdOutput) Listen() error {
	// THESE BOTH DO THE SAME THING, WHICHEVER ONE IS COMMENTED IS NOT THE ONE I'M USING RIGHT NOW

	// for line := range x.input {
	//  fmt.Println(line)
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
