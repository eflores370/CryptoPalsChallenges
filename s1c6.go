package main

import (
	"os"
	"sort"
	"strings"
)
import "fmt"
import "bufio"
import "encoding/base64"

type scores struct {
	scoreResult float32
	rawBytes    []byte
}

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

func findKeysize(byteArr []byte) {

	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		defer recoverloop()
		fmt.Println("Keysize:", KEYSIZE)
		// fmt.Print(byteArr[:KEYSIZE])
		// fmt.Println(byteArr[KEYSIZE:KEYSIZE*2])
		fmt.Println("Hamming Distance:", (hammingDist(byteArr[:KEYSIZE], byteArr[KEYSIZE:KEYSIZE*2]))/KEYSIZE)
	}
}

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

// Convert large byte array into chunks of smaller byte array
func breakCipherBlocks(byteArr []byte, Keysize int) (cipherBlocks [][]byte) {

	for i := 0; i < len(byteArr); i += Keysize {

		smallArry := make([]byte, 0)

		for j := 0; j < Keysize; j++ {

			if i+j < len(byteArr) {
				smallArry = append(smallArry, byteArr[i+j])
				defer recoverloop()
			}
		}
		cipherBlocks = append(cipherBlocks, smallArry)
	}

	return cipherBlocks
}

func transposeBlocks(Arr [][]byte, Keysize int) (Arr2 [][]byte) {

	for i := 0; i < Keysize; i++ {
		Arr2 = append(Arr2,make([]byte, 0))
	}

	for i := range Arr{
		for j := range Arr[i] {
			Arr2[j%Keysize] = append(Arr2[j%Keysize], Arr[i][j])
		}
	}
	return Arr2
}

func bruteforce(ByteArr[] byte) (scoreList[]scores) {


	// Brute force every character
	for i := 0; i < 255; i++ {
		tmpByteArr := make([]byte, len(ByteArr))
		copy(tmpByteArr, ByteArr)
		result := XOR(tmpByteArr, byte(i))

		scoreList = append(scoreList, score(result))
	}
	return scoreList
}

func score(rawBytes []byte) scores {
	englishFreq := map[string]float32{
		"E": 12.70, "T": 9.06, "A": 8.17, "O": 7.51, "I": 6.97,
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

	//fmt.Println(string(rawBytes), totalScore)

	return scores{scoreResult: totalScore, rawBytes: rawBytes}
}

func XOR(rawBytes []byte, key byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key
	}
	return rawBytes
}

func main() {

	var lines string
	// var block [][]byte

	// s := []byte("this is a test")
	// s1 := []byte("wokka wokka!!!")
	// fmt.Println(hammingDist(s,s1))

	// Place lines from file into large byte array
	file := readFile("files/6.txt")
	for i := range file {
		lines += file[i]
	}
	decoded, _ := base64.StdEncoding.DecodeString(lines)

	// Keysize is probably 5
	//findKeysize(decoded)
	
	chunckedArry := breakCipherBlocks(decoded, 5)
	ModifiedArray := transposeBlocks(chunckedArry, 5)

	for i := range ModifiedArray{
		//fmt.Println(ModifiedArray[i])
		scoreList := bruteforce(ModifiedArray[i])

		sort.Slice(scoreList, func(a, b int) bool { return scoreList[a].scoreResult < scoreList[b].scoreResult })

		for j := range scoreList {
			fmt.Println(scoreList[j].scoreResult)
			fmt.Println(scoreList[j].rawBytes)
			fmt.Println(string(scoreList[j].rawBytes))
		}
	}
}
