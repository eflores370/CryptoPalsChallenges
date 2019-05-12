package main

import "os"
import "fmt"
import "sort"
import "strings"
import "bufio"
import "encoding/hex"

type scores struct {
	scoreResult float32
	rawBytes    []byte
}

// Peform XOR to byte array with single byte key
func XOR(rawBytes []byte, key byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key
	}
	return rawBytes
}

// Perform Frequency Analysis and give a score to byte array
func score(rawBytes []byte) scores {
	englishFreq := map[string]float32{
		" ": 15, "E": 12.70, "T": 9.06, "A": 8.17, "O": 7.51, "I": 6.97,
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

	// if ! (totalScore <= 0) {
	// 	// fmt.Println(totalScore, string(rawBytes))
	// }

	return scores{scoreResult: totalScore, rawBytes: rawBytes}
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

func main() {

	// Input

	const inputString = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	scoreList := make([]scores, 0)
	lines := readFile("files/4.txt")

	for line := range lines {
		// Brute force every character
		for i := 0; i < 255; i++ {
			rawBytes, _ := hex.DecodeString(lines[line])
			result := XOR(rawBytes, byte(i))

			scoreList = append(scoreList, score(result))
		}
	}

	// fmt.Println(scoreList)

	// Sort

	sort.Slice(scoreList, func(i, j int) bool { return scoreList[i].scoreResult < scoreList[j].scoreResult })

	for i := range scoreList {
		fmt.Println(scoreList[i].scoreResult)
		fmt.Println(string(scoreList[i].rawBytes))
	}
}
