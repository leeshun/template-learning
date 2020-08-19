package main

import (
	"os"
	"text/template"
)

type source struct {
	Dist string
	Name string
}

func main() {
	tpl, err := template.ParseFiles("simple/simple.tpl")
	if err != nil {
		panic(err)
	}
	s := &source{
		Dist:"alice",
		Name: "bob",
	}
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}
