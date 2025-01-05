package main

import (
	"fmt"
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
	date := time.Now()
	fmt.Println("Command-line diary on", date.Day(), "of", date.Month(), date.Year(), ", at", date.Local().Hour(), ":", date.Local().Minute(), "hrs")
	var option string
	displayOptions()
	fmt.Scanln(&option)
	clearScreen()

	switch option {
	case "c":
		createEntry()
	case "u":
		updateEntry()
	case "r":
		readEntry()
	case "d":
		deleteEntry()
	default:
		fmt.Println("sorry, you chose an invalid option")
	}
	fmt.Println("Thanks for using the command-line diary. Bye!")
}
