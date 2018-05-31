package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println("WE ARE SO COOL!!! 🎉")
	statsInfo, err := os.Stdin.Stat()

	if err != nil {
		log.Fatal("OH NOES THIS IS SO BROKEN!!")
	}

	fmt.Printf("%#v", statsInfo)

	if statsInfo.Mode()&os.ModeNamedPipe == 0 {
		log.Fatal("TOO MANY THINGS OR NOT ENOUGH THINGS BUT THINGS ARE WRONG!")
	}
}
