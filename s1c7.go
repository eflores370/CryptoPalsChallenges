package main

import (
	"bufio"
	"crypto/aes"
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

func decryptAES128ECB(ciphertext, key []byte) (plaintext []byte) {
	cipher,_ := aes.NewCipher(key)
	plaintext = make([]byte, len(ciphertext))
	size := 16

	for blockStart,blockEnd:= 0, size; blockStart < len(ciphertext); blockStart, blockEnd = blockStart+size, blockEnd+size {
		cipher.Decrypt(plaintext[blockStart:blockEnd], ciphertext[blockStart:blockEnd])
	}

	return plaintext
}

func main() {

	key := []byte("YELLOW SUBMARINE")
	decoded := readFile("files/7.txt")

	plaintext := decryptAES128ECB(decoded, key)

	fmt.Println(string(plaintext))

}
