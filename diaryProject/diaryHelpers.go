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

// Creates a new file with filename format DDMMYYYY_HHMMhrs_entry.md, which contains diary entry
func createEntry() {
	date := time.Now()
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

	//what went well today
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
	// what didn't go well
	fmt.Println(bad)
	_, _ = file.WriteString(bad)
	file.WriteString("\n")
	badEntry, _ := reader.ReadString('\n')
	_, badWriteErr := file.WriteString(badEntry)
	if badWriteErr != nil {
		log.Fatalf("error when writing to file: %v", err)
	}
	file.WriteString("\n")
	// what could be improved today
	fmt.Println(improve)
	_, _ = file.WriteString(improve)
	file.WriteString("\n")
	improveEntry, _ := reader.ReadString('\n')
	_, improveWriteErr := file.WriteString(improveEntry)
	if improveWriteErr != nil {
		log.Fatalf("error when writing to file: %v", err)
	}
	clearScreen()

}

// Updates a diary entry file
func updateEntry() {
	// open and update an existing entry
	fmt.Println("Enter the filename of the entry:")
	filename := getFilename()
	reader := bufio.NewReader(os.Stdin)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	// get user io for entry.
	fmt.Println("Updated diary entry:")
	entryText, _ := reader.ReadString('\n')
	_, writeErr := file.WriteString(entryText)
	if writeErr != nil {
		log.Fatalf("error when writing to file: %v", err)
	}
	fmt.Println("Entry updated successfully!")
	clearScreen()
}

// Reads and displays diary entry in command line
func readEntry() {
	// read existing entry
	fmt.Println("Enter the filename of the entry:")
	filename := getFilename()
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
	fmt.Printf("%s has been deleted", filename)
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
