package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("JhnLrs")

type MyClaims struct {
	userid int `json:"user_id"`
	jwt.StandardClaims
}

func GenToken(id int) (string,error)  {
	claim := MyClaims{
		userid: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix()+7*24*60*60,
			Issuer: "JhnLrs",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	//加密
	ss,err :=token.SignedString(jwtKey)
	if err!=nil{
		return "",err
	}
	return ss,nil
	
}

func ParseToken(ss string) (*jwt.Token,*MyClaims,error) {
	claims := &MyClaims{}
	token,err := jwt.ParseWithClaims(ss,claims,func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token,claims,err
}