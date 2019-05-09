package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

// ReadFile
// Parse a file and returns the contents
// Input: File Path (String)
// Output: Array of Strings ([]string)
func readFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// ToBin
// Convert a byte array to a binary string
// Input: Byte Array ([]Byte)
// Output: Binary (String)

func ToBin(byteArr []byte) (binary string) {

	for i := range byteArr {
		binary += fmt.Sprintf("%08b", byteArr[i])
	}

	return binary

}

// HammingDist
// Returns the Hamming distance of two strings
// Input: Byte Array1 ([]byte), Byte Array2 ([]byte)
// Output: Distance (int)
func hammingDist(block1, block2 []byte) (distance float32) {

	binBlock1 := ToBin(block1)
	binBlock2 := ToBin(block2)

	for i := range binBlock1 {
		if binBlock1[i] != binBlock2[i] {
			distance++
		}
	}

	return distance

}

// FindKeySize
func findKeySize(initial, final int, byteArr []byte) (keysize int) {

	var shortestDist float32 = 255
	for KEYSIZE := initial; KEYSIZE <= final; KEYSIZE++ {
		//defer recoverloop()

		currentDist := float32(hammingDist(byteArr[:KEYSIZE], byteArr[KEYSIZE:KEYSIZE*2])) / float32(KEYSIZE)
		if currentDist < shortestDist {
			shortestDist = currentDist
			keysize = KEYSIZE
			//fmt.Println("Keysize:", KEYSIZE, "Hamming Distance", shortestDist)
		}
	}

	return keysize
}

// Given a large array, split the array into smaller array with the maximum size of Keysize
// Input: Large byte arr ([]byte) & Max length of array (int)
// Output: A Large array containing smaller arrays ([][]byte)
func breakCipherBlocks(byteArr []byte, Keysize int) (cipherBlocks [][]byte) {

	for i := 0; i < len(byteArr); i += Keysize {

		smallArry := make([]byte, 0)

		for j := 0; j < Keysize; j++ {

			if i+j < len(byteArr) {
				smallArry = append(smallArry, byteArr[i+j])
				//defer recoverloop()
			}
		}
		cipherBlocks = append(cipherBlocks, smallArry)
	}

	return cipherBlocks
}

func main() {

	var lines string

	// Read file and base 64 decode
	file := readFile("files/6.txt")
	for i := range file {
		lines += file[i]
	}
	decoded, _ := base64.StdEncoding.DecodeString(lines)

	keysize := findKeySize(2, 42, decoded)

	fmt.Println(keysize)
	fmt.Println(breakCipherBlocks(decoded, keysize))

}
