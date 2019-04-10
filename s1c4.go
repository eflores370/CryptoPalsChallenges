package main

import "fmt"
import "encoding/hex"


func XOR(rawBytes[]byte, key byte) []byte  {	
	for i := range rawBytes {
		rawBytes[i] ^= key
	}
	return rawBytes
}

func main(){
	
	// Input

	const inputString = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	

	for i := 0; i < 255; i++ {
		rawBytes, _ := hex.DecodeString(inputString)
		fmt.Println(string(rawBytes))
		result := XOR(rawBytes, byte(i))
		fmt.Println(string(result))
	}
}	
