package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math/rand"
)

func padding(unpaddedBytes []byte, totalLength int) (bytes []byte) {

	bytes = unpaddedBytes

	padValue := totalLength - (len(unpaddedBytes) % totalLength)

	for i := 0; i < padValue; i++ {
		bytes = append(bytes, byte(padValue))
	}

	return bytes
}

func XOR(rawBytes []byte, key []byte) []byte {
	for i := range rawBytes {
		rawBytes[i] ^= key[i%len(key)]
	}
	return rawBytes
}


func AES_128_CBC_Encrypt(plaintext string, key []byte) (ciphertext []byte) {
	paddedPlainText := padding([]byte(plaintext), 16)
	ciphertext = make([]byte, len(paddedPlainText))

	cipher, _ := aes.NewCipher(key)
	//iv := generateRandomBytes(16)
	iv := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

	for blockstart := 0; blockstart < len(plaintext); blockstart += aes.BlockSize {
		blockend := blockstart + aes.BlockSize
		copy(paddedPlainText[blockstart:blockend],XOR(paddedPlainText[blockstart:blockend], iv))
		cipher.Encrypt(ciphertext[blockstart:blockend], paddedPlainText[blockstart:blockend])
		iv = ciphertext[blockstart:blockend]
	}
	return ciphertext

}

func generateRandomBytes(length int) (key []byte) {

	key = make([]byte, length)
	_, err := rand.Read(key)

	if err != nil {
		return nil
	}

	return key
}

// Byte-at-a-time ECB decryption (Simple)
func main(){
	//key := generateRandomBytes(16)
	key := []byte("YELLOW SUBMARINE")
	unknownString := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	myString := "AAAA"

	decodedString,_ := base64.StdEncoding.DecodeString(unknownString)

	combinedString := myString + string(decodedString)

	fmt.Println(base64.StdEncoding.EncodeToString(AES_128_CBC_Encrypt(combinedString, key)))
}