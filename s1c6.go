package main

import "fmt"

func hammingDist(string1, string2 string) (distance int) {

	for i := range string1 {
		if (string1[i] != string2[i]) {
			distance ++
		}
	}

	return distance

}


// Convert a text string to a binary string
// Input: Text (String)
// Output: Binary (String)

func stringToBin(string1 string) (binary string) {

	for i := range string1 {
		binary += fmt.Sprintf("%08b", byte(string1[i]))
	}

	return binary

}

func main() {
	
	// var KEYSIZE int;

	s := "this is a test"

	s1 := "wokka wokka!!!"

	fmt.Println(hammingDist(stringToBin(s),stringToBin(s1)))
}