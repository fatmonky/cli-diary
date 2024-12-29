package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

/* Command-line diary project
The purpose of this project is to create a small command-line diary project.
This allows the creation of single short entries from the command-line, when the program is run.

The program should
- give clear instructions for what to do, allowing the user to
	- create an entry
	- update an existing entry
	- retrieve an existing entry
	- delete an existing entry
- print the date and time of the entry
- save each diary entry as a separate markdown file

*/

func main() {

	// TODO: print instructions for what to do when program is run
	date := time.Now()
	fmt.Println("Command-line diary on", date.Day(), "of", date.Month(), date.Year(), ", at", date.Local().Hour(), ":", date.Local().Minute(), "hrs")

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
