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
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM table_name")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var data []byte
		err := rows.Scan(&data)
		if err != nil {
			panic(err.Error())
		}

		encryptedData, err := Encrypt(data, key)
		if err != nil {
			panic(err.Error())
		}

		writeToFile(currentDir+"/encrypted_data.bin", encryptedData)
	}
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
