package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"taskbuilder/internal/core/domain"
	mock_port "taskbuilder/internal/mock"
	"testing"
)

func fakeUser() *domain.User {
	return &domain.User{
		Name:     "any name",
		Contact:  "any contact",
		Password: "any password",
		Email:    "any@me.com",
	}
}

func Test_userHandler_Register(t *testing.T) {
	actual := fakeUser()
	userJSON := `{
			  "name": "any name",
			  "email": "any@me.com",
			  "password": "any password",
			  "contact": "any contact"
			}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_port.NewMockUserService(ctrl)
	svc.EXPECT().Create(gomock.Any()).Return(actual, nil)
	h := NewUserHandler(svc)

	if assert.NoError(t, h.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func Test_userHandler_Register_failed(t *testing.T) {
	userJSON := `{
			  "name": "john k",
			  "email": "john@me.com",
			  "password": "1234",
			  "contact": "38399238"
			}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_port.NewMockUserService(ctrl)
	svc.EXPECT().Create(gomock.Any()).Return(nil, errors.New("an error occurred"))
	h := NewUserHandler(svc)

	if assert.NoError(t, h.Register(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func Test_userHandler_Index(t *testing.T) {
	actual := &domain.Users{*fakeUser()}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_port.NewMockUserService(ctrl)
	svc.EXPECT().GetAll().Return(actual, nil)
	h := NewUserHandler(svc)

	// Assertions
	if assert.NoError(t, h.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}

func Test_userHandler_Index_failed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock_port.NewMockUserService(ctrl)
	svc.EXPECT().GetAll().Return(nil, errors.New("an error occurred"))
	h := NewUserHandler(svc)

	if assert.NoError(t, h.Index(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}
