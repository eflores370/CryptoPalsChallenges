package main

import "fmt"
import "strings"
import "encoding/hex"


func XOR(rawBytes[]byte, key byte) []byte  {	
	for i := range rawBytes {
		rawBytes[i] ^= key
	}
	return rawBytes
}

func score(rawBytes[]byte) float32 {
	englishFreq := map[string]float32{
		"E": 12.70, "T": 9.06, "A": 8.17, "O": 7.51, "I": 6.97,
		"N": 6.75, "S": 6.33, "H": 6.09, "R": 5.99, "D": 4.25, "L": 4.03, 
		"C": 2.78, "U": 2.76, "M": 2.41, "W": 2.36, "F": 2.23, "G": 2.02, 
		"Y": 1.97, "P": 1.93, "B": 1.29, "V": 0.98, "K": 0.77, "J": 0.15, 
	 	"X": 0.15, "Q": 0.10, "Z": 0.07}

	var totalScore float32
	for i := range rawBytes {
		points, exists := englishFreq[strings.ToUpper(string(rawBytes[i]))]
		if exists {
			totalScore += points
		} 
	}

	if ! (totalScore <= 0) {
		fmt.Println(totalScore, string(rawBytes))	
	}
	
	return totalScore
}

func main(){
	
	// Input

	const inputString = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	

	for i := 0; i < 255; i++ {
		rawBytes, _ := hex.DecodeString(inputString)
		// fmt.Println(string(rawBytes))
		result := XOR(rawBytes, byte(i))
		// fmt.Println(string(result))
		score(result)
	}
}	
