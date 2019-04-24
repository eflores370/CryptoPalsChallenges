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
		fmt.Println(i)
		tmpArray := make([]byte, 0)
		for j := 0; j < len(Arr); j++ {
			//fmt.Println(Arr[j][i])
			tmpArray = append(tmpArray, Arr[j][i])
			defer recoverloop()
		}
		Arr2 = append(Arr2, tmpArray)

	}

	return Arr2
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

	// fmt.Println(decoded)
	chunckedArry := breakCipherBlocks(decoded, 5)
	fmt.Println(transposeBlocks(chunckedArry, 5))

}
