package main

import "fmt"
import "encoding/hex"

// Fixed XOR

func main() {
	const s = "1c0111001f010100061a024b53535009181c"
	const k = "686974207468652062756c6c277320657965"

	decodedS, _ := hex.DecodeString(s)
	decodedK, _ := hex.DecodeString(k)

	for i := range decodedS {
		decodedS[i] ^= decodedK[i]
	}

	output := hex.EncodeToString(decodedS)

	fmt.Println(output)
}
