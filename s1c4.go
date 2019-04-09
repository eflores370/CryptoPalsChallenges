package main

import "fmt"
// import "encoding/hex"
// import "io/ioutil"
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



func main() {

	// data,_ := ioutil.ReadFile("4.txt")
	// fmt.Print(string(data))

	lines := readLines("4.txt")

	for i := range lines {
		fmt.Println(lines[i])
	}
	
}