package util

import (
	"testing"
	"time"

	"github.com/fengshux/blog2/backend/conf"
)

func TestGenerateJWT(t *testing.T) {

	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 3600,
		},
	})

	var id int64 = 1
	token, err := GenerateJWT(id)
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

	var id int64 = 1
	token, err := GenerateJWT(id)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
		return
	}
	userId, err := extractClaims(token)
	if userId == 0 || err != nil {
		t.Error("extractClaims error:", err)
		return
	}

	if userId != id {
		t.Error("extractClaims expect id:", id, " get:", userId)
	}
}

func TestTokenExpire(t *testing.T) {
	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 1,
		},
	})

	var id int64 = 1
	token, err := GenerateJWT(id)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
		return
	}
	time.Sleep(time.Second * 2)
	userId, err := extractClaims(token)
	if userId != 0 || err == nil {
		t.Error("token not expired")
		return
	}
}
