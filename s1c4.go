package main

import "fmt"
import "encoding/hex"
import "bufio"
import "os"

func readLines(path string) ([]string) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}


	return lines
}

// func score() {
// englishFreq := map[string]float32{
// 	"E": 12.70, "T": 9.06, "A": 8.17, "O": 7.51, "I": 6.97,
// 	"N": 6.75, "S": 6.33, "H": 6.09, "R": 5.99, "D": 4.25, "L": 4.03, 
// 	"C": 2.78, "U": 2.76, "M": 2.41, "W": 2.36, "F": 2.23, "G": 2.02, 
// 	"Y": 1.97, "P": 1.93, "B": 1.29, "V": 0.98, "K": 0.77, "J": 0.15, 
//  	"X": 0.15, "Q": 0.10, "Z": 0.07}
// }

func bruteForce(s string){
	for i := 0; i < 255; i++ {
		decoded, _ := hex.DecodeString(s)
		for j := range decoded {
			decoded[j] ^= byte(i)
		}
		fmt.Println(i, string(i), string(decoded))
	}
}


func main() {

	// data,_ := ioutil.ReadFile("4.txt")
	// fmt.Print(string(data))

	lines := readLines("files/4.txt")

	for i := range lines {
		// fmt.Println(lines[i])
		decoded,_ := hex.DecodeString(lines[i])
		fmt.Println(string(decoded))

		bruteForce(lines[i])
	}
	
}