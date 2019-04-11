package main

import "os"
import "fmt"
import "bufio"
import "encoding/base64"

// Returns the Hamming distance of two strings
// Input: Byte Array1 ([]byte), Byte Array2 ([]byte)
// Output: Distance (int)

func hammingDist(block1, block2 []byte) (distance int) {
	
	binBlock1 := ToBin(block1)
	binBlock2 := ToBin(block2)

	for i := range binBlock1 {
		if binBlock1[i] != binBlock2[i] {
			distance++
		}
	}

	return distance

}



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

func findKEYSIZE(byteArr []byte) {
	

	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		defer recoverloop()
		fmt.Println("Keysize:", KEYSIZE)
		fmt.Print(byteArr[:KEYSIZE])
		fmt.Println(byteArr[KEYSIZE:KEYSIZE*2])
		fmt.Println(hammingDist(byteArr[:KEYSIZE],byteArr[KEYSIZE:KEYSIZE*2]))

	}

}

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

	file := readFile("files/6.txt")

	var lines string

	for i := range file {
		lines += file[i]
	}

	fmt.Println(lines)
	fmt.Println(base64.StdEncoding.DecodeString(lines))
	

	decoded, _ := base64.StdEncoding.DecodeString(lines)

	findKEYSIZE(decoded)

	// fmt.Println("HELLO")


	fmt.Println(hammingDist(s,s1))
}