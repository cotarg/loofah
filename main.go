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
	statsInfo, err := os.Stdin.Stat()
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Fatal("OH NOES THIS IS SO BROKEN!!")
	}

	if statsInfo.Mode()&os.ModeNamedPipe == 0 {
		log.Fatal("TOO MANY THINGS OR NOT ENOUGH THINGS BUT THINGS ARE WRONG!")
	}

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
}
