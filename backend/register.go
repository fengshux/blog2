package backend

import (
	"net/http"

	"github.com/fengshux/blog2/backend/controller"
	"github.com/fengshux/blog2/backend/util"
	"github.com/gin-gonic/gin"
)

type Register struct {
	user *controller.User
}

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

func newRegister(
	user *controller.User,
) *Register {
	return &Register{
		user: user,
	}
}

func (reg *Register) Regist(r *gin.RouterGroup) {
	r.GET("/user", ConvertController(reg.user.PageList))
}
