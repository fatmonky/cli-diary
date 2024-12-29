package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	fmt.Println("Choose the following options:")
	fmt.Println(`'c' to create new entry at this date/time
'u' to update an existing entry
'r' to read an existing entry
'd' to delete an existing entry`)
	fmt.Scanln(&option)
	clearScreen()

	switch option {
	case "c":
		//create entry file
		dayTimeString := fmt.Sprintf("%02d%02d%2d_%02d%02dhrs", date.Day(), date.Month(), date.Year(), date.Local().Hour(), date.Minute())
		filename := fmt.Sprintf("%s_entry.md", dayTimeString)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error when creating file: %v", err)
		}
		defer file.Close()
		// write date of entry to file
		file.WriteString("Entry for " + dayTimeString)
		file.WriteString("\n")
		// get user io for entry.
		fmt.Println("Diary entry:")
		well := ">>>What went well today?"
		bad := ">>>What didn't go well today?"
		improve := ">>>What could be improved today?"
		fmt.Println(well)
		_, _ = file.WriteString(well)
		file.WriteString("\n")
		reader := bufio.NewReader(os.Stdin)
		entryText, _ := reader.ReadString('\n')
		_, writeErr := file.WriteString(entryText)
		if writeErr != nil {
			log.Fatalf("error when writing to file: %v", err)
		}
		file.WriteString("\n")
		fmt.Println(bad)
		_, _ = file.WriteString(bad)
		file.WriteString("\n")
		badEntry, _ := reader.ReadString('\n')
		_, badWriteErr := file.WriteString(badEntry)
		if badWriteErr != nil {
			log.Fatalf("error when writing to file: %v", err)
		}
		file.WriteString("\n")
		fmt.Println(improve)
		_, _ = file.WriteString(improve)
		file.WriteString("\n")
		improveEntry, _ := reader.ReadString('\n')
		_, improveWriteErr := file.WriteString(improveEntry)
		if improveWriteErr != nil {
			log.Fatalf("error when writing to file: %v", err)
		}
	case "u":
		// open and update an existing entry
		fmt.Println("Enter the filename of the entry:")
		reader := bufio.NewReader(os.Stdin)
		filename, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer file.Close()
		// get user io for entry.
		fmt.Println("Diary entry:")
		entryText, _ := reader.ReadString('\n')
		_, writeErr := file.WriteString(entryText)
		if writeErr != nil {
			log.Fatalf("error when writing to file: %v", err)
		}
		fmt.Println("Entry updated successfully!")

	case "r":
		// read existing entry
		fmt.Println("Enter the filename of the entry:")
		reader := bufio.NewReader(os.Stdin)
		filename, _ := reader.ReadString('\n')
		filename = strings.Trim(filename, "\n")
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer file.Close()
		// fmt.Println(file)
		io.Copy(os.Stdout, file)

	case "d":
		// delete entry
		fmt.Println("Enter the file you want to delete:")
		reader := bufio.NewReader(os.Stdin)
		filename, _ := reader.ReadString('\n')
		filename = strings.Trim(filename, "\n")
		err := os.Remove(filename)
		if err != nil {
			log.Printf("Error in deleting file: %v", err)
		}
		fmt.Printf("%s has been deleted", filename)
	default:
		fmt.Println("sorry, you chose an invalid option")
	}
	clearScreen()
	fmt.Println("Thanks for using the command-line diary. Bye!")

}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Unix-like (Linux, macOS)
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
