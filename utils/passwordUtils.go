package utils

import (
	"io"
	"fmt"
	"bytes"
	"crypto/md5"
	"crypto/aes"
	"crypto/cipher"
)

const key = "\xc3\xd1w8\xaa\x9eCTk\xd4`^\n\xf3\xcc\xbf\xb4\xcbN\xbdx\xa9 d"


var keyByte = []byte(key)


func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src) % blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func makeAes(pwd string) (string, error) {
	pwdByte := []byte(pwd)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData := PKCS5Padding(pwdByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return string(crypted[:]), nil
}


func verifyAes(crypted string) (string, error) {
	cryptedByte := []byte(crypted)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = PKCS5UnPadding(origData)
	return string(origData[:]), nil
}


func makeMd5(crypted string) string {
	w := md5.New()
	io.WriteString(w, key)
	io.WriteString(w, crypted)
	// TODO  w.Sum(nil)将w的hash转成[]byte格式
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}


func MakePwd(pwd string) string {
	crypted, _ := makeAes(pwd)
	md5str := makeMd5(crypted)
	return md5str
}


func VerifyPwd(pwd string, md5str string) bool {
	b := md5str == MakePwd(pwd)
	return b
}
