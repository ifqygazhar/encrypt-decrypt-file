package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("program decryption v1.0")

	fmt.Print("masukan key encryption :")
	var inputKey string
	fmt.Scanln(&inputKey)

	fmt.Print("masukan nama file :")
	var inputFile string
	fmt.Scanln(&inputFile)

	cipherKey, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("error ioutil.ReadFIle")
		return
	}

	c, err := aes.NewCipher([]byte(inputKey))
	if err != nil {
		log.Fatal("error aes.Newchiper")
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal("error cipher.NewGCM")
		return
	}

	nonce := gcm.NonceSize()
	if len(cipherKey) < nonce {
		log.Fatal("error")
		return
	}

	nonceSize, chipherKey := cipherKey[:nonce], cipherKey[nonce:]
	plaintext, err := gcm.Open(nil, nonceSize, chipherKey, nil)
	if err != nil {
		log.Fatal("error cipher.NewGCM")
		return
	}

	fmt.Println(string(plaintext))

}
