package ruisUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Md5String(data string) string {
	sct := md5.New()
	sct.Write([]byte(data))
	md5Data := sct.Sum(nil)
	return hex.EncodeToString(md5Data)
}

func Md5(data []byte) string {
	sct := md5.New()
	sct.Write(data)
	md5Data := sct.Sum(nil)
	return hex.EncodeToString(md5Data)
}
func Sha1String(data string) string {
	sct := sha1.New()
	sct.Write([]byte(data))
	return hex.EncodeToString(sct.Sum(nil))
}
func Sha1(data []byte) string {
	sct := sha1.New()
	sct.Write(data)
	return hex.EncodeToString(sct.Sum(nil))
}
func Sha256String(data string) string {
	sct := sha256.New()
	sct.Write([]byte(data))
	return hex.EncodeToString(sct.Sum(nil))
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
