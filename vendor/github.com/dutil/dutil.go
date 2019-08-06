package dutil

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GetStrGUID 生成GUID
func GetStrGUID() string {
	u, _ := uuid.NewV1()
	str := u.String()
	str = strings.Replace(str, "-", "", -1)
	return StrMd5([]byte(str))
}

// StrMd5 16进制MD5
func StrMd5(bt []byte) string {
	h := md5.New()
	h.Write(bt[:])
	return fmt.Sprintf("%x", h.Sum(nil))
}

// ByteMd5 byte类型MD5
func ByteMd5(bt []byte) []byte {
	h := md5.New()
	h.Write(bt[:])
	return h.Sum(nil)
}

// SafeRandom 生成随机数
func SafeRandom(nums int) []byte {
	r := make([]byte, nums)
	if _, err := io.ReadFull(rand.Reader, r); err != nil {
		panic(err.Error())
	}
	return r
}
