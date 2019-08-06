package lib

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"

	"github.com/dgrijalva/jwt-go"
)

func NewMD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenToken() string {

	// nodeAndModeHex := hex.EncodeToString([]byte(fmt.Sprintf("%d.%d", node, 1)))
	// guid := dutil.GetStrGUID()
	// randomHex := hex.EncodeToString(genRand(5))
	// token := fmt.Sprintf("g1%s.%s%s", nodeAndModeHex, guid, randomHex)
	// return token

	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 60),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte("jwt test"))
	if err != nil {
		beego.Error("token error: ", err.Error())
		return ""
	}
	return str
}

func CheckToken(token string) error {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte("jwt test"), nil
	})
	return err
}
