package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 有效期
const JWTTimeExpireDuration = time.Minute * 1

// 加盐密钥
var secret = []byte("ZyRyJoZzz2ygTE1M8qOW0XFGcSulFRo40Fk871lYQp810RTW")

// 签发人
const Issuer = "xxcheng"

func Test_GenerateToken(t *testing.T) {
	mc := &MyClaims{
		UserID:   12345678,
		Username: "jpc",
		StandardClaims: jwt.StandardClaims{
			//使用UUID生成的一个，用于保证唯一
			Id: "d1414649-5a85-4a10-b1c5-befc1ac005d3",
			//设置过期时间
			ExpiresAt: time.Now().Add(JWTTimeExpireDuration).Unix(),
			//签发人
			Issuer: Issuer,
		},
	}
	//选择头部信息和载荷
	//此时是未加密的
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	//打印头部
	fmt.Printf("%T,%#v\n", token.Header, token.Header)
	//打印载荷
	fmt.Printf("%T,%#v\n", token.Claims, token.Claims)
	tokenStr, _ := token.SignedString(secret)
	fmt.Println("------")
	fmt.Println(tokenStr)
}

func Test_ParseToken(t *testing.T) {
	mc := new(MyClaims)
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NTY3OCwidXNlcm5hbWUiOiJqcGMiLCJleHAiOjE2OTExNTA4NzEsImp0aSI6ImQxNDE0NjQ5LTVhODUtNGExMC1iMWM1LWJlZmMxYWMwMDVkMyIsImlzcyI6Inh4Y2hlbmcifQ.bXZsPGy53yd-B2diEfejOkbWps8a87nJci8Qi45uuJw"
	_, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorExpired:
			fmt.Println("过期~~~")

		}
		return
	}
	fmt.Println(mc)
}
