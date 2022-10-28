package ruisUtil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

/*
生产私钥公钥的方法
openssl genrsa -out private.pem
openssl rsa -in private.pem -pubout -out public.pem
**/

func RsaSign(conts []byte, privatekeybts []byte) ([]byte, error) {
	// bts, err := ioutil.ReadFile("private.key")
	block, _ := pem.Decode(privatekeybts)
	if block == nil {
		return nil, fmt.Errorf("pem decode err")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("非法私钥文件，检查私钥文件")
	}

	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(conts))
	if err != nil {
		return nil, err
	}
	hashed := h.Sum(nil)
	signbts, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	//一般签名需要转一下base64
	return signbts, nil
}
func RsaVerify(conts []byte, signbts []byte, publickeybts []byte) error {
	// bts, err := ioutil.ReadFile("public.key")
	block, _ := pem.Decode(publickeybts)
	if block == nil {
		return fmt.Errorf("pem decode err")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("非法公钥文件，检查公钥文件")
	}

	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(conts))
	if err != nil {
		return err
	}
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, signbts)
}
