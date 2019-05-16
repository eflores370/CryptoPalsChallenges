package main

import (
	"fmt"
)

// Implement PKCS#7 padding

func padding(unpaddedText string, totalLength int) {

	bytes := []byte(unpaddedText)
	fmt.Println(bytes)

	padValue := totalLength - (len(unpaddedText) % totalLength)
	fmt.Println(padValue)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	fmt.Println(bytes)
}

func main() {
	const s = "YELLOW SUBMARINE"
	padding(s, 20)
	fmt.Println(string(s))
}
