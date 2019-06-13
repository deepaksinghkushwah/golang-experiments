package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var password = "this is secured password"

func main() {
	/*str := "Hello World"
	bin := []byte(str)
	fmt.Println(bin)

	binToStr := string(bin)
	fmt.Println(binToStr)*/

	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("Hello World"), password)
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decrypt(ciphertext, password)
	fmt.Printf("Decrypted: %s\n", plaintext)

	encryptFile("sample.dat", []byte("Hello World, this is in file"), password)
	fmt.Println(string(decryptFile("sample.dat", password)))
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, password string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(password)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText

}

func decrypt(data []byte, password string) []byte {
	key := []byte(createHash(password))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, password string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, password))
}

func decryptFile(filename string, password string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, password)
}
