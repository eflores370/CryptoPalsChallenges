package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
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


func AES_128_ECB_Encrypt(plaintext, key []byte) (ciphertext []byte){
	paddedPlainText := padding([]byte(plaintext), aes.BlockSize)
	ciphertext = make([]byte, len(paddedPlainText))

	cipher, _ := aes.NewCipher(key)
	for blockstart := 0; blockstart < len(paddedPlainText); blockstart += aes.BlockSize {
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

func detectBlockSize(ciphertext, key []byte) (length int) {
	currentBlockSize := len(AES_128_ECB_Encrypt(ciphertext, key))
	for i := 0; i < 50; i++ {
		stringA := strings.Repeat("A", i)
		tmpArray := make([]byte, 0)
		tmpArray = append(tmpArray,  stringA...)
		tmpArray = append(tmpArray,  ciphertext...)
		if currentBlockSize != len(AES_128_ECB_Encrypt(tmpArray, key)){
			return len(AES_128_ECB_Encrypt(tmpArray, key)) - currentBlockSize
		}

	}
	return 0
}

func detectECB(cipherText, key []byte, blocksize int) bool {

	counter := 0
	cipherBlocks := make([][]byte, 0)

	for blockStart := 0; blockStart < len(cipherText); blockStart += blocksize {
		blockEnd := blockStart + blocksize
		if len(cipherBlocks) == 0 {
			cipherBlocks = append(cipherBlocks, cipherText[blockStart:blockEnd])
			continue
		}
		for i := range cipherBlocks {
			if reflect.DeepEqual(cipherBlocks[i], cipherText[blockStart:blockEnd]) {
				counter++
			} else {
				cipherBlocks = append(cipherBlocks, cipherText[blockStart:blockEnd])
			}
		}
	}

	if counter > 0 {
		return true
	} else {
		return false
	}
}

func encryptionOracle(decodedString, key []byte, blocksize int) (plaintext string) {
	lookupTable := make(map[int][]byte)

	for i := 0; i < 255; i++ {
		controlledString := strings.Repeat("A", blocksize - 1)
		controlledString += string(i)
		encryptedString := AES_128_ECB_Encrypt([]byte(controlledString), key)
		lookupTable[i] = encryptedString
	}

	//for blockStart := 0; blockStart < len(ciphertext); blockStart += blocksize {
	//	blockEnd := blockStart + blocksize
	//	currentBlock := ciphertext[blockStart:blockEnd]
	//	for keyA, value := range lookupTable {
	//		fmt.Println(currentBlock,value[0:16], string(keyA))
	//		if reflect.DeepEqual(currentBlock,value[0:16]) {
	//			fmt.Print("!!!!!" + string(key))
	//		}
	//	}
	//}
	fmt.Println(lookupTable)
	fmt.Println(decodedString)
	for i := range decodedString {
		controlledString := strings.Repeat("A", blocksize - 1)
		controlledString += string(decodedString[i])
		encryptedString := AES_128_ECB_Encrypt([]byte(controlledString), key)
		for index, value := range lookupTable {
			//fmt.Println(encryptedString[0:blocksize], value[0:blocksize])
			if reflect.DeepEqual(encryptedString[0:blocksize], value[0:blocksize]) {
				plaintext += string(index)
			}
		}
	}

	return plaintext

}

// Byte-at-a-time ECB decryption (Simple)
func main(){
	//key := generateRandomBytes(16)
	key := []byte("YELLOW SUBMARINE")
	unknownString := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	//unknownString := "UEFTU1dPUkRERUFEQkVFRlBBU1NXT1JEREVBREJFRUY="

	// Decode string and append it to controlled string
	decodedString,_ := base64.StdEncoding.DecodeString(unknownString)

	// Encryption Setup
	cipherText := AES_128_ECB_Encrypt(decodedString, key)
	//fmt.Println(base64.StdEncoding.EncodeToString(cipherText))
	//fmt.Println(cipherText)

	// Detect Blocksize of the encryption
	blockSize := detectBlockSize(cipherText, key)
	//fmt.Println(blockSize)

	//fmt.Println(detectECB(cipherText, key, blockSize))

	plaintext := encryptionOracle(decodedString, key, blockSize)
	fmt.Println(plaintext)
}