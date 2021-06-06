package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/service"
	"taskbuilder/internal/server/rest/dto"
	"taskbuilder/internal/types"
)

type authHandler struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthHandler(authService service.AuthService, jwtService service.JwtService) *authHandler {
	return &authHandler{authService, jwtService}
}

func (hdl *authHandler) SignInWithEmailAndPassword(context echo.Context) error {
	var req dto.SignInCredential
	err := context.Bind(&req)
	if err != nil {
		return context.JSONPretty(http.StatusBadRequest, map[string]interface{}{"message": err.Error()}, "")
	}

	user, err := hdl.authService.SignInWithEmailAndPassword(req)
	if err != nil {
		return context.JSONPretty(http.StatusUnauthorized, map[string]interface{}{"message": err.Error()}, "")
	}
	payload := types.JwtPayload{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role.Name,
	}
	token, err := hdl.jwtService.GenerateToken(payload)
	if err != nil {
		fmt.Println(err.Error())
		return context.JSONPretty(http.StatusInternalServerError, map[string]interface{}{"message": "something went wrong"}, "")
	}
	response := dto.UserLoginResponse{
		AccessToken: token,
	}
	return context.JSONPretty(http.StatusOK, response, "")
}
