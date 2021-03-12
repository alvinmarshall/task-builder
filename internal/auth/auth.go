package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type JwtPayload struct {
	Id    string
	Email string
	Role  string
}

type JwtWrapper struct {
	Secret    string
	Issuer    string
	ExpiresAt int64
}

type JwtClaim struct {
	Id    string
	Email string
	Role  string
	jwt.StandardClaims
}

func GetJWTConfig(secret string) middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtClaim{},
		SigningKey: secret,
	}
}

func (j *JwtWrapper) GenerateToken(payload JwtPayload) (token string, err error) {
	claim := &JwtClaim{
		Email: payload.Email,
		Role:  payload.Role,
		Id:    payload.Id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.Issuer,
			Id:        payload.Id,
			Subject:   payload.Email,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.ExpiresAt)).Unix(),
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
