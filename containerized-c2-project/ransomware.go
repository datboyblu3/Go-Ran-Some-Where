package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func Decrypt(ciphertext []byte), key []byte ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func writeToFile(filename string, content []byte) error {
	tmpFilePath := filename + ".tmp"

	//Write content to a temp file
	err := os.WriteFile(tmpFilePath, content, 0644)
	if err != nil {
		return err
	}

	return os.NewFile()
}

func doEncrypt(currentDir string, key []byte) {

}

func doDecrypt(currentDir string, key []byte) {

}

func main() {
	key := []byte("PutYourKeyEncryptionKeyHere1234567890")

	//Get the current working directory
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	doEncrypt(currentDir, key)
	doDecrypt(currentDir, key)
}
