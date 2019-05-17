package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	random "math/rand"
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

func encryption_oracle(plaintext, key []byte) (ciphertext []byte) {

	random.Seed(time.Now().UTC().UnixNano())
	if random.Intn(100)%2 == 0 {
		//ECB
		ciphertext = ECBEncryption(plaintext, key)
	} else {
		//CBC
		ciphertext = CBCEncryption(plaintext, key)
	}

	return ciphertext
}

func main() {

	Plaintext := []byte("12345678123456781234567812345678")

	key := generateRandomBytes(16)
	encrypted := encryption_oracle(Plaintext, key)
	fmt.Println(encrypted)
	fmt.Println(len(encrypted))
	fmt.Println(string(encrypted))

}
