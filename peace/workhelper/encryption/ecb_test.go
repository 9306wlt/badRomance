package encryption

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "fmt"
    "strings"
    "testing"
)

func Test_ttt(t *testing.T) {
    /*
     *src 要加密的字符串
     *key 用来加密的密钥 密钥长度可以是128bit、192bit、256bit中的任意一个
     *16位key对应128bit
     */
    src := "&type=2&UA=Chrome&fpqqlsh=LSKP0062202012280001"
    key := "e6169f9e95cf4446c21c11d4e4c61110"
    fmt.Printf("明文：%v\n",src)
    crypted := Aes_Encrypt(src, key)
    fmt.Println("加密密文")
    fmt.Println(base64.StdEncoding.EncodeToString(crypted))
    jm := Aes_Decrypt(crypted, []byte(key))
    fmt.Println("解密明文")
    fmt.Println(string(jm))
    //Base64URLDecode("39W7dWTd_SBOCM8UbnG6qA")
}

func Base64URLDecode(data string) ([]byte, error) {
    var missing = (4 - len(data)%4) % 4
    data += strings.Repeat("=", missing)
    res, err := base64.URLEncoding.DecodeString(data)
    fmt.Println("  decodebase64urlsafe is :", string(res), err)
    return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
    // Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
    bytearr := base64.StdEncoding.EncodeToString(source)
    safeurl := strings.Replace(string(bytearr), "/", "_", -1)
    safeurl = strings.Replace(safeurl, "+", "-", -1)
    safeurl = strings.Replace(safeurl, "=", "", -1)
    return safeurl
}

func Aes_Decrypt(crypted, key []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println("err is:", err)
    }
    blockMode := NewECBDecrypter(block)
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PK_CS5UnPadding(origData)
    //fmt.Println("source is :", origData, string(origData))
    return origData
}

func Aes_Encrypt(src, key string) []byte {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        fmt.Println("key error1", err)
    }
    if src == "" {
        fmt.Println("plain content empty")
    }
    ecb := NewECBEncrypter(block)
    content := []byte(src)
    content = PK_CS5Padding(content, block.BlockSize())
    crypted := make([]byte, len(content))
    ecb.CryptBlocks(crypted, content)
    //// 普通base64编码加密 区别于urlsafe base64
    //fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))
    //
    //fmt.Println("base64UrlSafe result:", Base64UrlSafeEncode(crypted))
    return crypted
}

func PK_CS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PK_CS5UnPadding(origData []byte) []byte {
    length := len(origData)
    // 去掉最后一个字节 unpadding 次
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

type ecb struct {
    b         cipher.Block
    blockSize int
}

func newECB(b cipher.Block) *ecb {
    return &ecb{
        b:         b,
        blockSize: b.BlockSize(),
    }
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Encrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Decrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}
