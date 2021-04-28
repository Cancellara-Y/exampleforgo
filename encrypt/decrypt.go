package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func Decrypt(encodingContent, encodingAesKey string) (string, error) {
	base64DeContent, err := base64.StdEncoding.DecodeString(encodingContent)
	if err != nil {
		return "", err
	}

	deContent := AesDecryptCBC(base64DeContent, []byte(encodingAesKey))
	return string(deContent), nil
}

func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 分组秘钥
	blockSize := block.BlockSize()
	fmt.Println(blockSize)

	// 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	fmt.Println(len(encrypted))
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return nil
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
