package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(value, keyString string) string {
	// menguraikan kunci dan mengubah bentuk text encipt
	key, _ := hex.DecodeString(keyString)
	plainText := []byte(value)

	// membuat block cipher dari kunci
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// membuat GCM dengan counter mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// membuat nonce dari GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	fmt.Println(nonce)

	// Membuat enkripsi dari GCM seal
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	return fmt.Sprintf("%x", cipherText)
}

func Decrypt(encryptString, keyString string) string {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptString)

	// membuat block cipher dari kunci
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// membuat GCM dengan counter mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()
	fmt.Println(nonceSize)

	//Extract the nonce from the encrypted data
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plainText)
}

func GetRandKey() string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	// menencode key (bytes) menjadi string agar tidak terlihat
	key := hex.EncodeToString(bytes)
	return key
}
