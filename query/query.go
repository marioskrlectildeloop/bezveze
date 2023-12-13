package query

import (
	"bufio"
	"csv/columns"
	"csv/fs"
	"fmt"
	"io"
	"strings"
)

var UnexpectedError error = fmt.Errorf("Unexpected error happened")
var NothingFound error = fmt.Errorf("Nothing found")

func Value(column string) ([]string, error) {
	r := bufio.NewReader(fs.File())
	foundValues := make([]string, 0)
	for {

		b, _, err := r.ReadLine()
		if err == io.EOF {
			return foundValues, NothingFound
		}

		if err != nil {
			return nil, err
		}

		c := string(b)

		values := strings.Split(c, ",")
		foundValues = append(foundValues, values[columns.Index(column)])
	}
}
