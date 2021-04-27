package main

import (
	"encoding/hex"
	"github.com/forgoer/openssl"
	"fmt"
)

const mKey = "9763LB9SC541254FCE8ES83D02AD2526"


func encryptRemoteId(content string) (string, error) {
	if content == "" {
		return content, nil
	}
	key := []byte(mKey)
	data := []byte(content)

	a, err := openssl.AesECBEncrypt(data, key, "PKCS5")
	if err != nil {
		return "", err
	}

	hex_string_data := hex.EncodeToString(a)
	return hex_string_data, nil
}

func decryptRemoteId(encryptContent string) (string, error) {
	if encryptContent == "" {
		return encryptContent, nil
	}

	key := []byte(mKey)
	data, err := hex.DecodeString(encryptContent)
	if err != nil {
		return "",err
	}

	a, err := openssl.AesECBDecrypt(data, key, "PKCS5")
	if err != nil {
		return "",err
	}
	return string(a), nil
}

func main() {
	fmt.Println(encryptRemoteId("A2E619B7D0EB0DD9E3379DD65058EFFC"))
	fmt.Println(decryptRemoteId("A2E619B7D0EB0DD9E3379DD65058EFFC"))

	fmt.Println(encryptRemoteId("61AAB9615ED0D84E669412A8DDA11580901A2DDBF0758730A19ADA8F0B6DD9F8"))
	fmt.Println(decryptRemoteId("61AAB9615ED0D84E669412A8DDA11580901A2DDBF0758730A19ADA8F0B6DD9F8"))


}