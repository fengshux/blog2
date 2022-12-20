package util

import (
	"testing"
	"time"

	"github.com/fengshux/blog2/backend/conf"
	"github.com/fengshux/blog2/backend/model"
)

func TestGenerateJWT(t *testing.T) {

	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 3600,
		},
	})

	user := &model.User{ID: 1, Role: "general"}
	token, err := GenerateJWT(user)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
	}
}

func TestExtractClaims(t *testing.T) {
	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 3600,
		},
	})

	user := model.User{ID: 1, Role: "general"}
	token, err := GenerateJWT(&user)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
		return
	}
	claim, err := extractClaims(token)
	if claim.ID == 0 || err != nil {
		t.Error("extractClaims error:", err)
		return
	}

	if claim.ID != 1 || claim.Role != "general" {
		t.Error("extractClaims expect :", user, " get:", *claim)
	}
}

func TestTokenExpire(t *testing.T) {
	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 1,
		},
	})

	user := model.User{ID: 1, Role: "general"}
	token, err := GenerateJWT(&user)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
		return
	}
	time.Sleep(time.Second * 2)
	claim, err := extractClaims(token)
	if claim != nil || err == nil {
		t.Error("token not expired")
		return
	}
}
