package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	//"strings"
)

// Byte-at-a-time ECB decryption (Simple)

func generateRandomBytes(length int) (key []byte) {

	key = make([]byte, length)
	_, err := rand.Read(key)

	if err != nil {
		return nil
	}

	return key
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

func padding(unpaddedBytes []byte, totalLength int) (bytes []byte) {

	bytes = unpaddedBytes

	padValue := totalLength - (len(unpaddedBytes) % totalLength)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	return bytes
}

func oracle(input string, rawBytes, key []byte) []byte {

	bytes := []byte(input)
	bytes = append(bytes, rawBytes...)
	encryptedString := ECBEncryption(padding(bytes, 16), key)

	return encryptedString

}

func discoverBlockSize(decoded, key []byte) int {

	currentSize := 0
	previousSize := 0
	counter := 0
	blocksize := 0

	for i := 0; i < 50; i++ {
		input := strings.Repeat("A", i)
		encryptedString := oracle(input, decoded, key)
		fmt.Println(len(encryptedString))
		currentSize = len(encryptedString)
		if i != 0 {
			if currentSize != previousSize {
				counter++
				if counter == 2 {
					break
				}
			}
			if counter == 1 {
				blocksize++
			}
		}
		previousSize = currentSize
	}
	return blocksize

}

func main() {

	key := generateRandomBytes(16)
	const secret = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	decoded, _ := base64.StdEncoding.DecodeString(secret)

	fmt.Print("blocksize ", discoverBlockSize(decoded, key))

	//s := ECBEncryption(padding(decoded, 16),key)

}
