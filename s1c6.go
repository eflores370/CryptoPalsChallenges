package main

import "fmt"

// Returns the Hamming distance of two strings
// Input: Binary String 1 (String), Binary String 2 (String)
// Output: Distance (int)

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

func recoverloop() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func findKEYSIZE(s string) {
	

	for KEYSIZE := 2; KEYSIZE < 40; KEYSIZE++ {
		defer recoverloop()
		fmt.Println("\n\n",KEYSIZE)
		fmt.Println(stringToBin(s[:KEYSIZE]))
		fmt.Println(stringToBin(s[KEYSIZE:KEYSIZE*2]))
		fmt.Println(hammingDist(stringToBin(s[:KEYSIZE]),stringToBin(s[KEYSIZE:KEYSIZE*2]))/KEYSIZE)

	}

}

func main() {

	s := "this is a test"
	// s1 := "wokka wokka!!!"

	// fmt.Println(hammingDist(stringToBin(s),stringToBin(s1)))

	findKEYSIZE(s)

	fmt.Println("HELLO")
}