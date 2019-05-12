package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

// ReadFile
// Parse a file and returns the contents
// Input: File Path (String)
// Output: Array of Strings ([]string)
func readFile(path string) (decoded []byte) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	var output = ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := range lines {
		output += lines[i]
	}
	decoded, _ = base64.StdEncoding.DecodeString(output)

	return decoded
}

func main() {

	const key = "YELLOW SUBMARINE"

	fmt.Println(string(readFile("files/7.txt")))

}
