package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

// Implement CBC mode

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

func XOR(rawBytes []byte, key []byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key[i%len(key)]
	}
	return rawBytes
}

func decryptCBC(rawBytes, iv, key []byte) (plaintext []byte) {
	cipher, _ := aes.NewCipher(key)
	plaintext = make([]byte, len(rawBytes))
	size := 16

	for blockStart, blockEnd := 0, size; blockStart < len(rawBytes); blockStart, blockEnd = blockStart+size, blockEnd+size {
		nextIV := make([]byte, 16)
		copy(nextIV, rawBytes[blockStart:blockEnd])
		cipher.Decrypt(plaintext[blockStart:blockEnd], rawBytes[blockStart:blockEnd])
		copy(plaintext[blockStart:blockEnd], XOR(plaintext[blockStart:blockEnd], iv))
		iv = nextIV
	}
	return plaintext
}

func main() {
	decoded := readFile("files/10.txt")
	iv := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	key := []byte("YELLOW SUBMARINE")

	plaintext := decryptCBC(decoded, iv, key)
	fmt.Println(string(plaintext))

}
