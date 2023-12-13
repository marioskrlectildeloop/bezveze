package parser

import (
	columns2 "csv/columns"
	"fmt"
	"strings"
)

var EmptyStatement error = fmt.Errorf("Empty statement are not allowed")
var InvalidToken error = fmt.Errorf("Invalid token")
var ColumnNotExists error = fmt.Errorf("Invalid column. Column does not exist")

type Tokens struct {
	Column string
}

func Parse(stmt string) (Tokens, error) {
	if stmt == "" {
		return Tokens{}, EmptyStatement
	}

	tokens := strings.Split(stmt, " ")
	if len(tokens) != 4 {
		return Tokens{}, InvalidToken
	}

	sel := tokens[0]
	if strings.ToLower(sel) != "select" {
		return Tokens{}, InvalidToken
	}

	columns := tokens[1]
	if strings.ToLower(columns) != "*" {
		return Tokens{}, InvalidToken
	}

	from := tokens[2]
	if strings.ToLower(from) != "from" {
		return Tokens{}, InvalidToken
	}

	c := tokens[3]
	if !columns2.Exists(c) {
		return Tokens{}, ColumnNotExists
	}

	return Tokens{Column: c}, nil
}
