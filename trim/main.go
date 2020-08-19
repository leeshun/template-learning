package main

import (
	"os"
	"text/template"
)

type pair struct {
	First int
	Second int
}

func main() {
	tpl, err := template.ParseFiles("trim/trim.tpl")
	if err != nil {
		panic(err)
	}
	p := &pair{
		First:  10,
		Second: 2,
	}
	err = tpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
