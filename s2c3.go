package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	random "math/rand"
	"reflect"
	"strings"
	"time"
)

// An ECB/CBC detection oracle

func generateRandomBytes(length int) (key []byte) {

	key = make([]byte, length)
	_, err := rand.Read(key)

	if err != nil {
		return nil
	}

	return key
}

func XOR(rawBytes []byte, key []byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key[i%len(key)]
	}
	return rawBytes
}

func ECBEncryption(plaintext, key []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	cipher, _ := aes.NewCipher(key)
	size := 16

	for blockStart, blockEnd := 0, size; blockStart < len(plaintext); blockStart, blockEnd = blockStart+size, blockEnd+size {
		cipher.Encrypt(ciphertext[blockStart:blockEnd], plaintext[blockStart:blockEnd])
	}

	return ciphertext
}

func CBCEncryption(plaintext, key []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	cipher, _ := aes.NewCipher(key)
	iv := generateRandomBytes(16)
	size := 16

	for blockStart, blockEnd := 0, size; blockStart < len(plaintext); blockStart, blockEnd = blockStart+size, blockEnd+size {
		copy(plaintext[blockStart:blockEnd], XOR(plaintext[blockStart:blockEnd], iv))
		cipher.Encrypt(ciphertext[blockStart:blockEnd], plaintext[blockStart:blockEnd])
		copy(iv, ciphertext[blockStart:blockEnd])

	}

	return ciphertext
}

func padding(unpaddedBytes []byte, totalLength int) (bytes []byte) {

	bytes = unpaddedBytes

	padValue := totalLength - (len(unpaddedBytes) % totalLength)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	return bytes
}

func detectECB(ciphertext []byte) (counter int) {
	size := 16

	cipherBlocks := make([][]byte, 0)

	for blockStart, blockEnd := 0, size; blockStart < len(ciphertext); blockStart, blockEnd = blockStart+size, blockEnd+size {
		if len(cipherBlocks) == 0 {
			cipherBlocks = append(cipherBlocks, ciphertext[blockStart:blockEnd])
			continue
		}
		for i := range cipherBlocks {
			if reflect.DeepEqual(cipherBlocks[i], ciphertext[blockStart:blockEnd]) {
				counter++
			} else {
				cipherBlocks = append(cipherBlocks, ciphertext[blockStart:blockEnd])
			}
		}
	}
	return counter
}

func encryption_oracle(plaintext, key []byte) (ciphertext []byte) {

	random.Seed(time.Now().UTC().UnixNano())

	// Append 5 random bytes before & after
	tmpArray := generateRandomBytes(random.Intn(11-5) + 5)
	plaintext = append(tmpArray, plaintext...)
	plaintext = append(plaintext, generateRandomBytes(random.Intn(11-5)+5)...)

	// Add padding
	paddedText := padding(plaintext, 16)

	if random.Intn(100)%2 == 0 {
		//ECB
		ciphertext = ECBEncryption(paddedText, key)
	} else {
		//CBC
		ciphertext = CBCEncryption(paddedText, key)
	}

	if detectECB(ciphertext) > 0 {
		fmt.Println("ECB Detected")
		return ciphertext
	} else {
		fmt.Println("CBC Detected")
		return ciphertext
	}
}

func main() {

	input := strings.Repeat("A", 100)

	key := generateRandomBytes(16)
	fmt.Println(encryption_oracle([]byte(input), key))

}
