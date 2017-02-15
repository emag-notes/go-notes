package main

import (
	"os"
	"log"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}
