package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type UserHandler interface {
	Signup(echo.Context) error
	Login(echo.Context) error
	Restricted(echo.Context) error
	GetCurrentUser(echo.Context) error
	ShowUser(echo.Context) error
	ShowUsers(echo.Context) error
	HelloUser(echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (uh *userHandler) Restricted(c echo.Context) error {
	//トークンからユーザーをサーチ
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Name := claims["name"].(string)
	return c.String(http.StatusOK, Name)
}
func (uh *userHandler) HelloUser(c echo.Context) error {
	zap.S().Info("User API Hello")
	return c.String(http.StatusOK, "Hello, World!")
}

func (uh *userHandler) Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid Email or Password",
		}
	}

	createdUser, err := uh.userUsecase.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token := jwt.New(jwt.SigningMethodHS256)

	//トークンにユーザーの情報をセット
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = createdUser.Email
	claims["name"] = createdUser.Name
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
func (uh *userHandler) Login(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	loginUser, err := uh.userUsecase.FindByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if loginUser.ID == 0 || user.Password != loginUser.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid email or password",
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = loginUser.Email
	claims["name"] = loginUser.Name
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}

// GetCurrentUser の説明
func (uh *userHandler) GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	mail := claims["email"].(string)
	currentUser, err := uh.userUsecase.FindByEmail(mail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, currentUser)
}

// ShowUser の説明
func (uh *userHandler) ShowUser(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("user_id"))
	user, err := uh.userUsecase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (uh *userHandler) ShowUsers(c echo.Context) error {
	users, err := uh.userUsecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
