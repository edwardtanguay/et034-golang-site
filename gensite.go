// go script
package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the file name
	fileName := "example.txt"

	// Create or open the file for writing, create it if it doesn't exist, truncate if it does
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write content to the file
	content := "Hello, this is some text written to the file!"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and content written successfully.")
}
