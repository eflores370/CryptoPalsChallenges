package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
)

// Detect AES in ECB mode


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


func detectECB(ciphertext []byte) (counter int) {
	size := 16

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
		fmt.Println(hex.EncodeToString(decoded), detectECB(decoded))
	}
}
