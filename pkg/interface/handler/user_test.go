package handler

import (
	"testing"

	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func Test_userHandler_Restricted(t *testing.T) {
	type fields struct {
		userUsecase usecase.UserUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				userUsecase: tt.fields.userUsecase,
			}
			if err := uh.Restricted(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userHandler.Restricted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userHandler_Signup(t *testing.T) {
	type fields struct {
		userUsecase usecase.UserUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				userUsecase: tt.fields.userUsecase,
			}
			if err := uh.Signup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userHandler.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userHandler_Login(t *testing.T) {
	type fields struct {
		userUsecase usecase.UserUsecase
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userHandler{
				userUsecase: tt.fields.userUsecase,
			}
			if err := uh.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
