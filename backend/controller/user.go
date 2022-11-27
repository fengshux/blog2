package controller

import (
	"net/http"
	"strconv"

	"github.com/fengshux/blog2/backend/model"
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

	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}
	intSize, err := strconv.Atoi(size)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	opts := model.SQLOption{
		Limit:  intSize,
		Offset: (intPage - 1) * intSize,
	}

	users, err := u.userService.List(ctx, &opts)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	count, err := u.userService.Count(ctx)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return model.PageResponse[model.User]{
		List:  users,
		Total: count,
	}, nil
}

func (u *User) Create(ctx *gin.Context) (interface{}, util.HttpError) {

	user := model.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	_, err = u.userService.Create(ctx, &user)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return user, nil
}
