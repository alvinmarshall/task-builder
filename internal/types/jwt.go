package types

import (
	"github.com/dgrijalva/jwt-go"
	"taskbuilder/internal/config"
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

func NewJwtWrapper(config *config.Config) *JwtWrapper {
	return &JwtWrapper{
		Secret:    config.Jwt.Secret,
		Issuer:    config.Jwt.Issuer,
		ExpiresAt: config.Jwt.Expires,
	}
}
