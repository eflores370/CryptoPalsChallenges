package main

import "fmt"
import "encoding/hex"
import "encoding/base64"

// Convert Hex to Base64

func main() {
	
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decoded, _ := hex.DecodeString(s)

	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))
	fmt.Println(encoded)
}