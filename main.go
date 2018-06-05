package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

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

	if os.ModeNamedPipe == 0 {
		log.Fatal("TOO MANY THINGS OR NOT ENOUGH THINGS BUT THINGS ARE WRONG!")
	}
}
