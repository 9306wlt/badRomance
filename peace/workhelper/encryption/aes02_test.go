package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"testing"
)

func AES_CBC_Decrypt_Byts(cipherText, key []byte) (string, error) {
	defer aeccatch()

	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//指定初始化向量IV,和加密的一致
	iv := []byte("vMUUiTjo4UzreO4n")
	//blockSize := block.BlockSize()
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//删除填充
	//plainText = pKCS5Trimming(plainText)
	return string(plainText), nil
}

func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func aeccatch() error {
	if err := recover(); err != nil {
		fmt.Printf("aes error: %v\n", err)
		return fmt.Errorf("aes error: %v", err)
	}
	return nil
	//if err := recover(); err != nil {
	//	return
	//}
}

func Test_aes02(t *testing.T) {
	//*uz^3T_!M6v@8#aK
	key := "*uz^3T_!M6v@8#aK"
	plaintext := "N2pnQy4XRGvk6E0p5aFQmA=="
	fmt.Println("base64后密文：", plaintext)
	//iv := make([]byte,16)
	//crytedByte,_ := base64.StdEncoding.DecodeString(plaintext)
	mw, err := AES_CBC_Decrypt_Byts([]byte(plaintext), []byte(key))
	if err != nil {
		panic(err)
	}
	fmt.Println(mw)
}
