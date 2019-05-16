package main

import (
	"fmt"
)

// Implement PKCS#7 padding

func padding(unpaddedText string, totalLength int) (bytes []byte) {

	bytes = []byte(unpaddedText)
	fmt.Println(bytes)

	padValue := totalLength - (len(unpaddedText) % totalLength)
	fmt.Println(padValue)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	return bytes
}

func main() {
	const s = "YELLOW SUBMARINE"
	paddedText := padding(s, 20)
	fmt.Println(string(paddedText))
}
