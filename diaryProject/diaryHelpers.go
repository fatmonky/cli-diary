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

const (
	donePrompt    = ">>>What did you accomplish today? 💪"
	wellPrompt    = ">>>What went well today? 😃"
	badPrompt     = ">>>What didn't go well today? 😢"
	improvePrompt = ">>>What could be improved today? 🛠️"
)

func createFilename() string {
	//create entry file
	date := time.Now()
	dayTimeString := fmt.Sprintf("%02d%02d%02d_%02d%02dhrs", date.Year(), date.Month(), date.Day(), date.Local().Hour(), date.Minute())
	filename := fmt.Sprintf("%s_entry.md", dayTimeString)
	return filename
}

// Writes diary entry, when creating the file, using the respective prompt
func writeEntry(filename *os.File, prompt string) {
	fmt.Println(prompt)
	_, _ = filename.WriteString(prompt)
	filename.WriteString("\n")
	entryText := getUserInput()
	_, writeErr := filename.WriteString(entryText)
	if writeErr != nil {
		log.Fatalf("error when writing to file: %v", writeErr)
	}
	filename.WriteString("\n")
}

// Creates a new diary entry with timestamped filename.
func createEntry() {
	clearScreen()
	filename := createFilename()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error when creating file: %v", err)
	}
	defer file.Close()
	// write date of entry to file
	file.WriteString("Entry for " + filename[:len(filename)-9])
	file.WriteString("\n")
	// get user io for entry.
	fmt.Println("Diary entry:")

	//what did you accomplish today
	writeEntry(file, donePrompt)
	//what went well today
	writeEntry(file, wellPrompt)
	// what didn't go well
	writeEntry(file, badPrompt)
	// what could be improved today
	writeEntry(file, improvePrompt)
	clearScreen()
	fmt.Println("Entry created successfully!")
	fmt.Println("========Entry=======")
	displayEntry(filename)
	fmt.Println("========End=======")
}

// Updates a diary entry file
func updateEntry() {
	clearScreen()
	// open and update an existing entry
	fmt.Println("Enter the filename of the entry:")
	filename := getFilename()
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	// get user io for entry.
	fmt.Println(">>>Addendum:")
	file.WriteString(">>>Addendum:\n")
	entryText := getUserInput()
	_, writeErr := file.WriteString(entryText)
	if writeErr != nil {
		log.Fatalf("error when writing to file: %v", err)
	}
	clearScreen()
	fmt.Println("Entry updated successfully!")
	fmt.Println("========Entry=======")
	displayEntry(filename)
	fmt.Println("========End=======")
}

// Reads and displays diary entry in command line
func readEntry() {
	// read existing entry
	fmt.Println("Enter the filename of the entry:")
	filename := getFilename()
	fmt.Println("========Entry=======")
	displayEntry(filename)
	fmt.Println("========End=======")
}

// Displays entry
func displayEntry(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

// Deletes a diary entry file
func deleteEntry() {
	// delete entry
	fmt.Println("Enter the file you want to delete:")
	filename := getFilename()
	err := os.Remove(filename)
	if err != nil {
		log.Printf("Error in deleting file: %v", err)
	}
	fmt.Printf("%s has been deleted\n", filename)
}

// Get user input
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

// Clears terminal screen
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

// Gets filename from commandline
func getFilename() string {
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	filename = strings.Trim(filename, "\n")
	return filename
}

// Displays program menu options, to create, update, read or delete existing entries.
func displayOptions() {
	fmt.Println("Choose the following options:")
	fmt.Println(`'c' to create new entry at this date/time
'u' to update an existing entry
'r' to read an existing entry
'd' to delete an existing entry`)
}
