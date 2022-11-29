package backend

import (
	"net/http"

	"github.com/fengshux/blog2/backend/controller"
	"github.com/fengshux/blog2/backend/util"
	"github.com/gin-gonic/gin"
)

type ControllerFunc func(*gin.Context) (interface{}, util.HttpError)

func ConvertController(f ControllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := f(ctx)
		if err != nil {
			ctx.JSON(err.Status(), gin.H{"msg": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

type Register struct {
	user *controller.User
	post *controller.Post
}

func newRegister(
	user *controller.User,
	post *controller.Post,
) *Register {
	return &Register{
		user: user,
		post: post,
	}
}

func (reg *Register) Regist(r *gin.RouterGroup) {
	r.GET("/user", ConvertController(reg.user.PageList))
	r.POST("/user", ConvertController(reg.user.Create))
	r.POST("/signin", ConvertController(reg.user.Signin))

	r.GET("/post", ConvertController(reg.post.PageList))
	r.POST("/post", util.HardAuth(), ConvertController(reg.post.Create))
}
