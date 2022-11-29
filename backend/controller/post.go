package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fengshux/blog2/backend/model"
	"github.com/fengshux/blog2/backend/service"
	"github.com/fengshux/blog2/backend/util"
	"github.com/gin-gonic/gin"
)

type Post struct {
	postService *service.Post
}

func NewPost(postService *service.Post) *Post {
	return &Post{
		postService: postService,
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
		Limit:  intSize,
		Offset: (intPage - 1) * intSize,
	}

	posts, err := p.postService.List(ctx, &opts)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	count, err := p.postService.Count(ctx)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return model.PageResponse[model.Post]{
		List:  posts,
		Total: count,
	}, nil
}

func (p *Post) Create(ctx *gin.Context) (interface{}, util.HttpError) {

	post := model.Post{}
	err := ctx.ShouldBind(&post)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	// TODO use ctx.userId
	post.UserId = 1

	_, err = p.postService.Create(ctx, &post)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return post, nil
}
