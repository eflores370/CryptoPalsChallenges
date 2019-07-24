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


//func AES_128_CBC_Encrypt(plaintext string, key []byte) (ciphertext []byte) {
//	paddedPlainText := padding([]byte(plaintext), aes.Blocksize)
//	ciphertext = make([]byte, len(paddedPlainText))
//
//	cipher, _ := aes.NewCipher(key)
//	//iv := generateRandomBytes(16)
//	iv := []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x10\x11\x12\x13\x14\x15")
//
//	for blockstart := 0; blockstart < len(plaintext); blockstart += aes.BlockSize {
//		blockend := blockstart + aes.BlockSize
//		copy(paddedPlainText[blockstart:blockend],XOR(paddedPlainText[blockstart:blockend], iv))
//		cipher.Encrypt(ciphertext[blockstart:blockend], paddedPlainText[blockstart:blockend])
//		iv = ciphertext[blockstart:blockend]
//	}
//	return ciphertext
//
//}

func AES_128_ECB_Encrypt(plaintext string, key []byte) (ciphertext []byte){
	paddedPlainText := padding([]byte(plaintext), aes.BlockSize)
	ciphertext = make([]byte, len(paddedPlainText))

	cipher, _ := aes.NewCipher(key)
	for blockstart := 0; blockstart < len(plaintext); blockstart += aes.BlockSize {
		blockend := blockstart + aes.BlockSize
		cipher.Encrypt(ciphertext[blockstart:blockend], paddedPlainText[blockstart:blockend])
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
	cipherText := AES_128_ECB_Encrypt(combinedString, key)
	fmt.Println(base64.StdEncoding.EncodeToString(cipherText))
}