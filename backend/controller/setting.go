package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fengshux/blog2/backend/model"
	"github.com/fengshux/blog2/backend/service"
	"github.com/fengshux/blog2/backend/util"
	"github.com/gin-gonic/gin"
)

type Setting struct {
	settingService *service.Setting
}

func NewSetting(settingService *service.Setting) *Setting {
	return &Setting{
		settingService: settingService,
	}
}

func (s *Setting) List(ctx *gin.Context) (interface{}, util.HttpError) {

	settings, err := s.settingService.List(ctx, nil)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return settings, nil
}

func (s *Setting) Info(ctx *gin.Context) (interface{}, util.HttpError) {

	key := ctx.Param("key")
	if key == "" {
		return nil, util.NewHttpError(http.StatusForbidden, fmt.Errorf("参数错误"))
	}

	setting, err := s.settingService.FindOne(ctx, key)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return setting, nil
}

func (s *Setting) Update(ctx *gin.Context) (interface{}, util.HttpError) {

	setting := model.Setting{}

	err := ctx.ShouldBind(&setting)
	if err != nil {
		log.Println(err)
		return nil, util.NewHttpError(http.StatusBadRequest, err)
	}

	if setting.Key == "" {
		return nil, util.NewHttpError(http.StatusBadRequest, fmt.Errorf("参数错误"))
	}

	dbSetting, err := s.settingService.FindOne(ctx, setting.Key)
	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	// TODO 以下逻辑需要用事务
	if dbSetting != nil {
		_, err = s.settingService.Update(ctx, setting.Key, setting.Data)
	} else {
		_, err = s.settingService.Create(ctx, &setting)
	}

	if err != nil {
		return nil, util.NewHttpError(http.StatusInternalServerError, err)
	}

	return setting, nil
}
