package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
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

	users, err := u.userService.List(ctx, model.SQLWhere{}, &opts)
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

	user := model.FullUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}
	user.Role = "general"

	if user.Password != "" {
		password := md5.Sum([]byte(user.Password))
		user.Password = hex.EncodeToString(password[0:])
	}

	_, err = u.userService.Create(ctx, &user)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return user, nil
}

func (u *User) Signin(ctx *gin.Context) (interface{}, util.HttpError) {

	body := model.FullUser{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	if body.Password == "" || body.UserName == "" {
		return nil, util.NewHttpError(http.StatusBadRequest, fmt.Errorf("用户名或密码为空"))
	}

	user, err := u.userService.FindOneFullUser(ctx, &model.User{UserName: body.UserName})
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return nil, util.NewHttpError(http.StatusBadRequest, fmt.Errorf("用户名或密码错误"))
	}

	password := md5.Sum([]byte(body.Password))
	md5Password := hex.EncodeToString(password[0:])

	if user.Password != md5Password {
		return nil, util.NewHttpError(http.StatusBadRequest, fmt.Errorf("用户名或密码错误"))
	}

	token, err := util.GenerateJWT(&user.User)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}
	ctx.Header("Authorization", token)

	return gin.H{"msg": "login success", "role": user.Role}, nil
}
