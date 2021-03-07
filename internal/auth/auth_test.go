package auth

import (
	"fmt"
	"testing"
)

func TestJwtWrapper_GenerateToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		Secret:    "secret",
		ExpiresAt: 2,
		Issuer:    "task-builder",
	}
	payload := JwtPayload{
		Id:    "1",
		Email: "email@me.com",
		Role:  "user",
	}
	token, err := jwtWrapper.GenerateToken(payload)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(token)

}

func TestJwtWrapper_ValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJFbWFpbCI6ImVtYWlsQG1lLmNvbSIsIlJvbGUiOiJ1c2VyIiwiZXhwIjoxNjE1MTU1Mzg1LCJqdGkiOiIxIiwiaWF0IjoxNjE1MTU1MjY1LCJpc3MiOiJ0YXNrLWJ1aWxkZXIiLCJzdWIiOiJlbWFpbEBtZS5jb20ifQ.ncBFNcPtL5tf2Zwrty_nnyzYa1ypwOjF-jkkGvlYZEs"
	jwtWrapper := JwtWrapper{
		Secret: "secret",
	}
	jwtClaim, err := jwtWrapper.ValidateToken(token)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(jwtClaim)
}
