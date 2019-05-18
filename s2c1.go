package main

import (
	"fmt"
)

// Implement PKCS#7 padding

func padding(unpaddedBytes []byte, totalLength int) (bytes []byte) {

	bytes = unpaddedBytes

	padValue := totalLength - (len(unpaddedBytes) % totalLength)
	fmt.Println(padValue)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	return bytes
}

func main() {
	s := []byte("YELLOW SUBMARINE")
	paddedText := padding(s, 20)
	fmt.Println(string(paddedText))
}
