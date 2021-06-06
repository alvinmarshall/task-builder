package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"taskbuilder/internal/types"
	"time"
)

type JwtService interface {
	GetJWTConfig() middleware.JWTConfig
	GenerateToken(payload types.JwtPayload) (token string, err error)
	ValidateToken(token string) (claim *types.JwtClaim, err error)
}

type jwtService struct {
	jwtWrapper *types.JwtWrapper
}

func (j jwtService) GetJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &types.JwtClaim{},
		SigningKey: []byte(j.jwtWrapper.Secret),
	}
}

func (j jwtService) GenerateToken(payload types.JwtPayload) (token string, err error) {
	claim := &types.JwtClaim{
		Email: payload.Email,
		Role:  payload.Role,
		Id:    payload.Id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.jwtWrapper.Issuer,
			Id:        payload.Id,
			Subject:   payload.Email,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.jwtWrapper.ExpiresAt)).Unix(),
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = withClaims.SignedString([]byte(j.jwtWrapper.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j jwtService) ValidateToken(token string) (claim *types.JwtClaim, err error) {
	claims := &types.JwtClaim{}
	parseWithClaims, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.jwtWrapper.Secret), nil
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

func NewJwtService(wrapper *types.JwtWrapper) JwtService {
	return &jwtService{wrapper}
}
