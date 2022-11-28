package util

import "testing"

func TestGenerateJWT(t *testing.T) {

	var id int64 = 1
	token, err := GenerateJWT(id)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("token is empty")
	}
}
