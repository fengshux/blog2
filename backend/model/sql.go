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
	Operator SQLOperator
	Value    interface{}
}

type SQLWhere []SQLCondition

type SQLOperator string

var (
	SQLOperator_EQ      SQLOperator = "="
	SQLOperator_GT      SQLOperator = ">"
	SQLOperator_LT      SQLOperator = "<"
	SQLOperator_GTE     SQLOperator = ">="
	SQLOperator_LTE     SQLOperator = "<="
	SQLOperator_NE      SQLOperator = "<>"
	SQLOperator_IN      SQLOperator = "IN"
	SQLOperator_LK      SQLOperator = "LIKE"
	SQLOperator_TSQUERY SQLOperator = "@@"
)

func (w SQLWhere) ToGormHere() (string, []interface{}) {
	statements := make([]string, len(w))
	params := make([]interface{}, len(w))

	for i, v := range w {
		switch v.Operator {
		case SQLOperator_TSQUERY:
			statements[i] = fmt.Sprintf("%s %s plainto_tsquery('jiebacfg', ?)", v.Name, v.Operator)
		default:
			statements[i] = fmt.Sprintf("%s %s ?", v.Name, v.Operator)
		}
		params[i] = v.Value
	}

	return strings.Join(statements, " AND "), params
}
