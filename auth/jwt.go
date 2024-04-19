package auth

import (
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	User interface{}
	jwt.RegisteredClaims
}

const signKey = "RNXE3U31ZS00C6CP3KT8OY028OK19K8Y"

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func GenerateToken(user interface{}, issue string, subject string, audience []string, expiresAt time.Time) (string, error) {
	claim := &MyCustomClaims{
		user,
		jwt.RegisteredClaims{
			Issuer:    issue,                                           // 签发者
			Subject:   subject,                                         // 签发对象
			Audience:  jwt.ClaimStrings(audience),                      //签发受众
			ExpiresAt: jwt.NewNumericDate(expiresAt),                   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // jwt ID, 类似于盐值
		},
	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// sign
	if tokenString, err := token.SignedString([]byte(signKey)); err != nil {
		return "", err
	} else {
		// fmt.Printf("tokenString %v \n", tokenString)
		return tokenString, nil
	}
}

func ParseToken(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
