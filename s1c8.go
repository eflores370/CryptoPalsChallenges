package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
)

func readFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

//func decryptAES128ECB(ciphertext, key []byte) (plaintext []byte) {
//	cipher,_ := aes.NewCipher(key)
//	plaintext = make([]byte, len(ciphertext))
//	size := 16
//
//	for blockStart,blockEnd:= 0, size; blockStart < len(ciphertext); blockStart, blockEnd = blockStart+size, blockEnd+size {
//		cipher.Decrypt(plaintext[blockStart:blockEnd], ciphertext[blockStart:blockEnd])
//	}
//
//	return plaintext
//}

func detectECB(ciphertext []byte) (counter int) {
	size := 16

	// Create an empty array
	// For every block in ciphertext, check if current block exists in array
	// If block doesn't exist, store block in array
	// Else add to counter

	cipherBlocks := make([][]byte, 0)

	for blockStart, blockEnd := 0, size; blockStart < len(ciphertext); blockStart, blockEnd = blockStart+size, blockEnd+size {
		if len(cipherBlocks) == 0 {
			cipherBlocks = append(cipherBlocks, ciphertext[blockStart:blockEnd])
			continue
		}
		for i := range cipherBlocks {
			if reflect.DeepEqual(cipherBlocks[i], ciphertext[blockStart:blockEnd]) {
				counter++
			} else {
				cipherBlocks = append(cipherBlocks, ciphertext[blockStart:blockEnd])
			}
		}
	}
	return counter
}

func main() {

	lines := readFile("files/8.txt")

	for i := range lines {
		decoded, _ := hex.DecodeString(lines[i])
		fmt.Println(decoded, detectECB(decoded))
	}
}
