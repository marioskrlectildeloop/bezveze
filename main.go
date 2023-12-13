package main

import (
	"csv/columns"
	"csv/fs"
	"csv/parser"
	"csv/query"
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := fs.Open("csv.csv"); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := fs.Close(); err != nil {
			fmt.Println(fmt.Sprintf("Shutdown produced errors: %s", err.Error()))
		}
	}()

	if err := columns.Init(); err != nil {
		log.Fatalln(err)
	}

	tokens, err := parser.Parse("SELECT * FROM Variable_category")
	if err != nil {
		fmt.Println(err)
	}

	values, err := query.Value(tokens.Column)
	if !errors.Is(err, query.NothingFound) {
		log.Fatalln(err)
	}

	fmt.Println(values)
}
