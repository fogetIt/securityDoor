package utils

import (
	"fmt"
	"time"
	"strconv"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken2(UserId string) string {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"iss": "sso",
		"jti": UserId,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix() + 100,
		"exp": int64(time.Now().Unix() + 7 * 24 * 60 * 60),
	}
	ss, err := token.SignedString(keyByte)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}


func GenerateToken(UserId string) string {
	claims := &jwt.StandardClaims{
		Id: UserId,
		NotBefore: int64(time.Now().Unix() + 10),
		ExpiresAt: int64(time.Now().Unix() + 7 * 24 * 60 * 60),
		Issuer: "sso",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(keyByte)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}


func VerifyToken(token string) uint {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return keyByte, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return 0
	}
	var jti = t.Claims.(jwt.MapClaims)["jti"]
	UserId, err := strconv.Atoi(jti.(string))
	return uint(UserId)
}
