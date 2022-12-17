package model

import (
	"fmt"
	"strings"
)

type SQLOption struct {
	Limit   int
	Offset  int
	OrderBy string
}

type SQLCondition struct {
	Name     string
	Operator string
	Value    interface{}
}

type SQLWhere []SQLCondition

func (w SQLWhere) ToGormHere() (string, []interface{}) {
	statements := make([]string, len(w))
	params := make([]interface{}, len(w))

	for i, v := range w {
		statements[i] = fmt.Sprintf("%s %s ?", v.Name, v.Operator)
		params[i] = v.Value
	}

	return strings.Join(statements, " AND "), params
}
