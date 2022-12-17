package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fengshux/blog2/backend/model"
	"github.com/fengshux/blog2/backend/service"
	"github.com/fengshux/blog2/backend/util"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type Post struct {
	postService *service.Post
	userService *service.User
}

func NewPost(postService *service.Post, userService *service.User) *Post {
	return &Post{
		postService: postService,
		userService: userService,
	}
}

func (p *Post) PageList(ctx *gin.Context) (interface{}, util.HttpError) {

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
		Limit:   intSize,
		Offset:  (intPage - 1) * intSize,
		OrderBy: "id desc",
	}

	posts, err := p.postService.List(ctx, &opts)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	count, err := p.postService.Count(ctx)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	userIds := lo.Map(posts, func(p model.Post, _ int) int64 {
		return p.UserId
	})

	users, err := p.userService.List(ctx, model.SQLWhere{{"id", "in", userIds}}, nil)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}
	userMap := lo.KeyBy(users, func(u model.User) int64 {
		return u.ID
	})

	postVOs := lo.Map(posts, func(p model.Post, _ int) model.PostVO {
		var user *model.User
		if u, ok := userMap[p.UserId]; ok {
			user = &u
		}

		return model.PostVO{
			Post: p,
			User: user,
		}
	})

	return model.PageResponse[model.PostVO]{
		List:  postVOs,
		Total: count,
	}, nil
}

func (p *Post) Info(ctx *gin.Context) (interface{}, util.HttpError) {

	strId := ctx.Param("id")
	if strId == "" {
		return nil, util.NewHttpError(http.StatusForbidden, fmt.Errorf("参数错误"))
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	post, err := p.postService.FindOne(ctx, &model.Post{ID: int64(id)})
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	user, err := p.userService.FindOne(ctx, &model.User{ID: post.UserId})
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return model.PostVO{Post: *post, User: user}, nil
}

func (p *Post) Create(ctx *gin.Context) (interface{}, util.HttpError) {

	loginUserId := ctx.GetInt64("userId")

	if loginUserId == 0 {
		return nil, util.NewHttpError(http.StatusUnauthorized, fmt.Errorf("请登录"))
	}

	post := model.Post{}
	err := ctx.ShouldBind(&post)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	post.UserId = loginUserId

	_, err = p.postService.Create(ctx, &post)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return post, nil
}

func (p *Post) Update(ctx *gin.Context) (interface{}, util.HttpError) {

	loginUserId := ctx.GetInt64("userId")

	if loginUserId == 0 {
		return nil, util.NewHttpError(http.StatusUnauthorized, fmt.Errorf("请登录"))
	}

	strId := ctx.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	post := model.Post{}
	err = ctx.ShouldBind(&post)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}
	// 不能更新ID
	post.ID = 0

	err = p.postService.Updates(ctx, model.SQLWhere{{"id", "=", id}}, &post)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return post, nil
}

func (p *Post) Delete(ctx *gin.Context) (interface{}, util.HttpError) {
	loginUserId := ctx.GetInt64("userId")

	if loginUserId == 0 {
		return nil, util.NewHttpError(http.StatusUnauthorized, fmt.Errorf("请登录"))
	}

	strId := ctx.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	err = p.postService.Delete(ctx, int64(id))
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return `{"msg":"删除成功"}`, nil
}
