package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nRunning...")

	in := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter filename: ")
	filename, _ := in.ReadString('\n')
	filename = strings.TrimSpace(filename)

	// fmt.Print("Enter content to add: ")
	// data, _ :=in.ReadString('\n')
	data := "Sentence 1\nSentence 2\nSentence 3\nSentence 4\nSentence chama 5\nSentence 6"

	createFile(filename)
	writeFile(data, filename)
	readData := readFile(filename)
	lines := countLines(readData)
	found := findWord("chama", readData)

	fmt.Println("\nNumber of Lines: ", lines)
	fmt.Println("\nWord found: ", found)

	writeToNewFile(readData)
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("\nFailed to create file: " + err.Error())
	}
	defer file.Close()
}

func writeFile(data string, filename string) {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// os.O_CREATE: Creates the file if it doesn’t exist.
	// os.O_WRONLY: Opens the file in write-only mode.
	// os.O_TRUNC: Truncates the file if it already exists.
	// 0644: Sets file permissions (read/write for owner, read-only for others).

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	length, err := file.WriteString(data)
	if err != nil {
		fmt.Println("\nError writing to file:", err)
	} else {
		fmt.Println("\nText written successfully!")
	}

	fmt.Printf("\nFile name: %s\n", file.Name())
	fmt.Printf("File length: %d bytes\n", length)
}

func readFile(filename string) string {
	fmt.Println("\nReading File...")

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("\nError reading file:", err)
		return string(err.Error())
	}

	fmt.Println("\n", string(data))
	return string(data)
}

func countLines(data string) int {
	return strings.Count(data, "\n") + 1
}

func findWord(findMe string, content string) string {
	words := strings.Fields(content)
	for _, word := range words {
		if word == findMe {
			return word
		}
	}
	return ""
}

func writeToNewFile(content string) {
	writeFile(content, "newFile.txt")
}
