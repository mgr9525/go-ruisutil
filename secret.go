package ruisUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func Md5(data []byte) string {
	md5 := md5.New()
	md5.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}
func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
func Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func AESEncrypt(src, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if iv == nil || len(iv) <= 0 {
		iv = make([]byte, block.BlockSize())
	}
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, iv)
	dist := make([]byte, len(src))
	blockmode.CryptBlocks(dist, src)
	return dist, nil
}

func AESDecrypt(src, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if iv == nil || len(iv) <= 0 {
		iv = make([]byte, block.BlockSize())
	}
	blockmode := cipher.NewCBCDecrypter(block, iv)
	dist := make([]byte, len(src))
	blockmode.CryptBlocks(dist, src)
	dist = unpadding(dist)
	return dist, nil
}
