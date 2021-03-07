package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtWrapper struct {
	Secret    string
	Issuer    string
	ExpiresAt int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(email string) (token string, err error) {
	claim := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.ExpiresAt)).Unix(),
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = withClaims.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JwtWrapper) ValidateToken(token string) (claim *JwtClaim, err error) {
	claims := &JwtClaim{}
	parseWithClaims, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Secret), nil
		})
	if err != nil {
		return nil, err
	}
	err = parseWithClaims.Claims.Valid()
	if err != nil {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}
