package util

import (
	"testing"

	"github.com/fengshux/blog2/backend/conf"
)

func TestGenerateJWT(t *testing.T) {

	conf.SetConf(&conf.Conf{
		Auth: conf.Auth{
			Secret:  "1234567",
			Expires: 10,
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
