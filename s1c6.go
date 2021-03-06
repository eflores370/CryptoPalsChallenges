package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Break repeating-key XOR

type scores struct {
	scoreResult float32
	rawBytes    []byte
	key         int
}

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

		tmpArr := breakCipherBlocks(byteArr, KEYSIZE)
		var total float32 = 0

		distances := make([]float32, 0)
		for i := range tmpArr {
			if i+2 < len(tmpArr) {
				distances = append(distances, hammingDist(tmpArr[i], tmpArr[i+1])/float32(KEYSIZE))
			}
		}

		for i := range distances {
			total += distances[i]
		}

		currentDist := float32(total / float32(len(distances)))

		if currentDist < shortestDist {
			shortestDist = currentDist
			keysize = KEYSIZE
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
			}
		}
		cipherBlocks = append(cipherBlocks, smallArry)
	}

	return cipherBlocks
}

// Given a 2D array, create a new array
func transposeBlocks(Arr [][]byte, Keysize int) (Arr2 [][]byte) {

	for i := 0; i < Keysize; i++ {
		Arr2 = append(Arr2, make([]byte, 0))
	}

	for i := range Arr {
		for j := range Arr[i] {
			Arr2[j%Keysize] = append(Arr2[j%Keysize], Arr[i][j])
		}
	}
	return Arr2
}

// XOR a raw byte array with a given key
// Input: Byte Array ([]byte) & Key ([]byte)
func XOR(rawBytes []byte, key byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key
	}
	return rawBytes
}

// Scoring function to determine if output is probable English
// Input: Array of bytes ([]Byte)
// Output: A score struct containing total points and a byte array ({int, []byte})
func score(rawBytes []byte, value int) scores {
	englishFreq := map[string]float32{
		" ": 15, "E": 12.70, "T": 9.06, "A": 8.17, "O": 7.51, "I": 6.97,
		"N": 6.75, "S": 6.33, "H": 6.09, "R": 5.99, "D": 4.25, "L": 4.03,
		"C": 2.78, "U": 2.76, "M": 2.41, "W": 2.36, "F": 2.23, "G": 2.02,
		"Y": 1.97, "P": 1.93, "B": 1.29, "V": 0.98, "K": 0.77, "J": 0.15,
		"X": 0.15, "Q": 0.10, "Z": 0.07}

	var totalScore float32

	totalScore = 0

	for i := range rawBytes {
		points, exists := englishFreq[strings.ToUpper(string(rawBytes[i]))]
		if exists {
			totalScore += points
		}
	}

	return scores{scoreResult: totalScore, rawBytes: rawBytes, key: value}
}

// Brute force the XOR Key for a given Byte array and return a an array of scores
// Input: Byte array ([]byte)
// Output: Array of byte arrays
func bruteForce(ByteArr []byte) (XORByteArr []scores) {

	// Brute force every character
	for i := 0; i < 255; i++ {
		tmpByteArr := make([]byte, len(ByteArr))
		copy(tmpByteArr, ByteArr)
		result := XOR(tmpByteArr, byte(i))
		XORByteArr = append(XORByteArr, score(result, i))

	}
	return XORByteArr
}

// Recovers a loop from Runtime exceptions

func recoverloop() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func rebuildString(finalArry []scores, Keysize int) (finalString []byte) {

	for j := 0; j < Keysize; j++ {
		for i := 0; i < len(finalArry[0].rawBytes); i++ {
			for _, v := range finalArry {
				rawBytes := v.rawBytes
				finalString = append(finalString, rawBytes[i])
				defer recoverloop()
			}
		}
	}

	return finalString
}

func main() {

	var lines string
	//scoreList := make([][]byte, 0)
	finalArray := make([]scores, 0)

	// Read file and base 64 decode
	file := readFile("files/6.txt")
	for i := range file {
		lines += file[i]
	}
	decoded, _ := base64.StdEncoding.DecodeString(lines)

	keysize := findKeySize(2, 42, decoded)

	splitArray := breakCipherBlocks(decoded, keysize)
	ModifiedArray := transposeBlocks(splitArray, keysize)

	for i := range ModifiedArray {
		tmpArray := bruteForce(ModifiedArray[i])
		sort.Slice(tmpArray, func(a, b int) bool { return tmpArray[a].scoreResult < tmpArray[b].scoreResult })
		finalArray = append(finalArray, tmpArray[len(tmpArray)-1])
	}

	fmt.Println(string(rebuildString(finalArray, keysize)))

	// Key
	fmt.Print("Key: ")
	for i := range finalArray {
		fmt.Print(string(finalArray[i].key))
	}

}
