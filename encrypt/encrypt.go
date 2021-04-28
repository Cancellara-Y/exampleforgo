package encrypt

import (
	"encoding/hex"
	"github.com/forgoer/openssl"
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
