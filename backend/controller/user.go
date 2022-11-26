package controller

import (
	"fmt"
	"net/http"

	"github.com/fengshux/blog2/backend/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	userService *service.User
}

func NewUser(user *service.User) *User {
	return &User{
		userService: user,
	}
}

func (u *User) PageList(ctx *gin.Context) {
	users, err := u.userService.List(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			fmt.Sprintf("{\"msg\": %s}", err),
		)
	}

	ctx.JSON(http.StatusOK, users)
}