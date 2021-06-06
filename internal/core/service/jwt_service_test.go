package service

import (
	"fmt"
	"taskbuilder/internal/types"
	"testing"
)

func TestJwtWrapper_GenerateToken(t *testing.T) {
	jwtWrapper := &types.JwtWrapper{
		Secret:    "secret",
		ExpiresAt: 2,
		Issuer:    "task-builder",
	}
	newJwtService := NewJwtService(jwtWrapper)
	payload := types.JwtPayload{
		Id:    "1",
		Email: "email@me.com",
		Role:  "user",
	}
	token, err := newJwtService.GenerateToken(payload)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(token)

}

func TestJwtWrapper_ValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJFbWFpbCI6ImVtYWlsQG1lLmNvbSIsIlJvbGUiOiJ1c2VyIiwiZXhwIjoxNjE1MTU1Mzg1LCJqdGkiOiIxIiwiaWF0IjoxNjE1MTU1MjY1LCJpc3MiOiJ0YXNrLWJ1aWxkZXIiLCJzdWIiOiJlbWFpbEBtZS5jb20ifQ.ncBFNcPtL5tf2Zwrty_nnyzYa1ypwOjF-jkkGvlYZEs"
	jwtWrapper := &types.JwtWrapper{
		Secret: "secret",
	}
	newJwtService := NewJwtService(jwtWrapper)
	jwtClaim, err := newJwtService.ValidateToken(token)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(jwtClaim)
}
