package main

import "fmt"
import "encoding/hex"

func main(){

	s1 := "Burning 'em, if you ain't quick and nimble"
	s2 := "\nI go crazy when I hear a cymbal"

	rawBytes1 := []byte(s1)
	rawBytes2 := []byte(s2)

	k := []byte("ICE")

	fmt.Println(hex.EncodeToString(rawBytes1))
	fmt.Println(hex.EncodeToString(rawBytes2))

	for i := range rawBytes1 {
		rawBytes1[i] ^= k[i%3]
	}

	for i := range rawBytes2 {
		rawBytes2[i] ^= k[i%3]
	}

	fmt.Println(hex.EncodeToString(rawBytes1))
	fmt.Println(hex.EncodeToString(rawBytes2))

}