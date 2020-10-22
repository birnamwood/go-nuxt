package handler

import (
	"net/http"
	"time"

	"github.com/birnamwood/go-nuxt/model"
	"github.com/birnamwood/go-nuxt/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Signup の説明
func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid Email or Password",
		}
	}
	user.Create()

	token := jwt.New(jwt.SigningMethodHS256)

	//トークンにユーザーの情報をセット
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//トークン生成
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}

//Login comment
func Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	user := repository.FindUserByEmail(u.Email)

	if user.ID == 0 || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid email or password",
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		//トークンからユーザーをサーチ
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		Name := claims["name"].(string)
		return c.String(http.StatusOK, Name)
	}
}
