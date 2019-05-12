package main

import "fmt"
import "encoding/hex"

// Implement repeating-key XOR

func main() {

	s := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	rawBytes1 := []byte(s)

	k := []byte("ICE")

	for i := range rawBytes1 {
		rawBytes1[i] ^= k[i%3]
	}

	fmt.Println(hex.EncodeToString(rawBytes1))

}
