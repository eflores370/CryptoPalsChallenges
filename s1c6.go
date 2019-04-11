package main

import "os"
import "fmt"
import "bufio"
// import "encoding/base64"

// Returns the Hamming distance of two strings
// Input: Binary String 1 (String), Binary String 2 (String)
// Output: Distance (int)

func hammingDist(block1, block2 []byte) (distance int) {
	
	blockA := ToBin(block1)
	blockB := ToBin(block2)

	for i := range blockA {
		if blockA[i] != blockB[i] {
			distance++
		}
	}

	return distance

}


// TO Remove
// func retrieveUMask(byteArr []byte) (umask string) {

// 	for i := range byteArr {
// 		for j := 2; j < 255; j *= 2 {
// 			fmt.Println(byteArr[i],j)
// 			fmt.Println(byteArr[i] & byte(j) != 0)
// 			if (byteArr[i] & byte(j) != 0){
// 				umask += "0"	
// 			} else {
// 				umask += "1"
// 			}
// 		}
// 	}
// 	fmt.Println(umask)

// 	return umask
// }


// Convert a byte array to a binary string
// Input: Byte Array ([]Byte)
// Output: Binary (String)

func ToBin(byteArr []byte) (binary string) {

	for i := range byteArr {
		binary += fmt.Sprintf("%08b", byteArr[i])
	}

	return binary

}

func recoverloop() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

// func findKEYSIZE(s string) {
	

// 	for KEYSIZE := 2; KEYSIZE < 40; KEYSIZE++ {
// 		defer recoverloop()
// 		fmt.Println(stringToBin(s[:KEYSIZE]))
// 		fmt.Println(stringToBin(s[KEYSIZE:KEYSIZE*2]))
// 		fmt.Println(hammingDist(stringToBin(s[:KEYSIZE]),stringToBin(s[KEYSIZE:KEYSIZE*2]))/KEYSIZE)

// 	}

// }

func readFile(path string) ([]string) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {

	s := []byte("this is a test")
	s1 := []byte("wokka wokka!!!")

	// file := readFile("files/6.txt")

	// var lines string

	// for i := range file {
	// 	lines += file[i]
	// }

	// fmt.Println(lines)
	// fmt.Println(base64.StdEncoding.DecodeString(lines))
	

	// fmt.Println(hammingDist(stringToBin(s),stringToBin(s1)))

	// a, _ := base64.StdEncoding.DecodeString(lines)

	// findKEYSIZE(s)

	// fmt.Println("HELLO")

	// hammingDist(byte)

	fmt.Println(s)	

	// fmt.Println(hammingDist(s,s1))

	fmt.Println(ToBin(s))

	fmt.Println(hammingDist(s,s1))
}