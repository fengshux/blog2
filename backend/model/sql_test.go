package model

import "testing"

func TestToGormWhere(t *testing.T) {
	where := SQLWhere{
		{
			"name",
			"=",
			"xiaoming",
		},
		{
			"id",
			"in",
			[]int64{1, 2},
		},
	}

	statement, params := where.ToGormHere()

	if statement != "name = ? AND id in ?" {
		t.Errorf("expect name = ? AND id in ? got %s", statement)
	}

	_, ok := params[1].([]int64)
	if params[0] != "xiaoming" || !ok {
		t.Errorf("params error %+v", params)
	}
}
