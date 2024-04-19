package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/dzeqkon/pkg/strings"
	"github.com/dzeqkon/pkg/symbol"
)

/*
signing a message
using: hmac sha256 + base64

	eg:
	  message = Pre_hash function comment
	  secretKey = E65791902180E9EF4510DB6A77F6EBAE

	return signed string = TO6uwdqz+31SIPkd4I+9NiZGmVH74dXi+Fd5X0EzzSQ=
*/
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

/*
md5 sign: "123" -> "202cb962ac59075b964b07152d234b70"
*/
func Md5Signer(message string) string {
	data := []byte(message)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GUID() string {
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	guid := Md5Signer(base64.URLEncoding.EncodeToString(b))
	return guid
}

/*
return eg: e7486845-9f24-c3d8-0db1-fe61e25c88a2
*/
func UUID() string {
	str := strings.NewString(GUID())
	builder := strings.NewStringBuilder()
	builder.Append(str.Substring(0, 8).ToString())
	builder.Append(symbol.HLINE).Append(str.Substring(8, 12).ToString())
	builder.Append(symbol.HLINE).Append(str.Substring(12, 16).ToString())
	builder.Append(symbol.HLINE).Append(str.Substring(16, 20).ToString())
	builder.Append(symbol.HLINE).Append(str.SubstringBegin(20).ToString())
	return builder.ToString()
}
