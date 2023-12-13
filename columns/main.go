package columns

import (
	"csv/fs"
	"csv/rd"
	"strings"
)

var columns []string

func Init() error {
	b, err := rd.Once(fs.File())

	if err != nil {
		return err
	}

	br := string(b)

	// no data validation, assume csv in correct format, who gives a fuck
	columns = strings.Split(br, ",")

	return nil
}

func All() []string {
	return columns
}

func Exists(c string) bool {
	for _, v := range columns {
		if v == c {
			return true
		}
	}

	return false
}

func Index(c string) int {
	for i, v := range columns {
		if v == c {
			return i
		}
	}

	return -1
}
