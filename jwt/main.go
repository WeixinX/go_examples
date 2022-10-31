package main

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// structureJWT 构造 Token
func structureJWT(uname, uid, issuer, subject, secret string, duration time.Duration) (string, error) {
	// 构造 payload
	t := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,                              // 签发人
		Subject:   subject,                             // 主题
		Audience:  jwt.ClaimStrings{uname},             // 受众
		ExpiresAt: jwt.NewNumericDate(t.Add(duration)), // 失效时间
		NotBefore: jwt.NewNumericDate(t),               // 生效时间
		IssuedAt:  jwt.NewNumericDate(t),               // 签发时间
		ID:        uid,                                 // 编号
	}
	// 计算 token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // head + payload
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err
}

// validateToken 解析并验证 Token
func validateToken(token, secret string) bool {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil && jwtToken != nil {
		if _, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
			return true
		}
	}
	return false
}

func main() {
	secret := "THISISATESTSECRET." // 密钥
	duration := 15 * time.Second   // token 有效期
	token, err := structureJWT("test", "-1", "sys", "login", secret, duration)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	ok := validateToken(token, secret)
	if !ok {
		fmt.Println("token 过期")
	} else {
		fmt.Println("token 正确")
	}
}
