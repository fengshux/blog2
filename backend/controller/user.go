package controller

import (
	"net/http"

	"github.com/fengshux/blog2/backend/service"
	"github.com/fengshux/blog2/backend/util"
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

func (u *User) PageList(ctx *gin.Context) (interface{}, util.HttpError) {
	users, err := u.userService.List(ctx)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return users, nil
}
