package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func AES_CBC_Decrypt_Byts04(cipherText, key []byte) (string, error) {
	defer aeccatch04()

	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//指定初始化向量IV,和加密的一致
	iv := []byte("vMUUiTjo4UzreO4n")

	blockSize := block.BlockSize()
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	//解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//删除填充
	//plainText = PKCS5UnPadding(plainText)
	return string(plainText), nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aeccatch04() error {
	if err := recover(); err != nil {
		fmt.Printf("aes error: %v\n", err)
		return fmt.Errorf("aes error: %v", err)
	}
	return nil
	//if err := recover(); err != nil {
	//	return
	//}
}

func Test_aes04(t *testing.T) {

	//plaintext := "N2pnQy4XRGvk6E0p5aFQmA=="
	//秘钥：*uz^3T_!M6v@8#aK

	//打印加密base64后密码
	key := "*uz^3T_!M6v@8#aK"
	plaintext := "N2pnQy4XRGvk6E0p5aFQmA=="
	fmt.Println("base64后密文：", plaintext)
	crytedByte, _ := base64.StdEncoding.DecodeString(plaintext)
	decryptCode, err := AES_CBC_Decrypt_Byts04(crytedByte, []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println("解密：", decryptCode)

}
