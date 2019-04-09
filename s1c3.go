package main

import "fmt"
import "encoding/hex"

// Single Byte XOR Cipher
// Return the original message that has been XOR with a single character

func main(){

	
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decoded, _ := hex.DecodeString(s)
	fmt.Println("Original:",string(decoded))

	// for i := 0; i<255; i++ {
	// 	key,_ := hex.DecodeString(string(i))
	// 	decoded[i] ^= key
	// 	fmt.Println(string(decoded))
	// }

	// for i := range decoded {
	// 	decoded, _ := hex.DecodeString(s)
	// 	for j := 0; j<255; j++ {
	// 		decoded[i] ^= byte(j)
	// 	}
	// 	fmt.Println(i, string(decoded))
	// }

	for i := 0; i < 255; i++ {
		decoded, _ := hex.DecodeString(s)
		for j := range decoded {
			decoded[j] ^= byte(i)
		}
		fmt.Println(i, string(i), string(decoded))
	}
	

}