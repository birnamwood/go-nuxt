package handler

import (
	"testing"

	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func Test_messageHandler_SendMessage(t *testing.T) {
	type fields struct {
		messageUsecase usecase.MessageUsecase
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
			mh := &messageHandler{
				messageUsecase: tt.fields.messageUsecase,
			}
			if err := mh.SendMessage(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("messageHandler.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
