package main

import (
	"io"
	"log"
	"os"
)

func main() {
	
	// opens file with filename
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	// fmt.Println(file)
	io.Copy(os.Stdout, file)

}
