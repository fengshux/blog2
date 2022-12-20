package backend

import (
	"log"
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
			log.Println(err)
			ctx.JSON(err.Status(), gin.H{"msg": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

type Register struct {
	user    *controller.User
	post    *controller.Post
	setting *controller.Setting
}

func newRegister(
	user *controller.User,
	post *controller.Post,
	setting *controller.Setting,
) *Register {
	return &Register{
		user:    user,
		post:    post,
		setting: setting,
	}
}

func (reg *Register) Regist(r *gin.RouterGroup) {
	r.GET("/user", ConvertController(reg.user.PageList))
	r.POST("/user", ConvertController(reg.user.Create))
	r.DELETE("/user/:id", util.AdminAuth(), ConvertController(reg.user.Delete))
	r.PATCH("/user/:id/password", util.AdminAuth(), ConvertController(reg.user.ChangePass))

	r.POST("/signin", ConvertController(reg.user.Signin))

	r.GET("/post", ConvertController(reg.post.PageList))
	r.POST("/post", util.HardAuth(), ConvertController(reg.post.Create))
	r.GET("/post/:id", ConvertController(reg.post.Info))
	r.PUT("/post/:id", util.HardAuth(), ConvertController(reg.post.Update))
	r.DELETE("/post/:id", util.HardAuth(), ConvertController(reg.post.Delete))

	r.GET("/setting", ConvertController(reg.setting.List))
	r.GET("/setting/:key", ConvertController(reg.setting.Info))
	r.PUT("/setting", util.AdminAuth(), ConvertController(reg.setting.Update))
}
