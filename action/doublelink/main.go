package main

import (
	"fmt"

	"github.com/leeshun/template-learning/util"
)

type Second struct {
	Value int
}

type First struct {
	Ref *Second
}

func main() {
	f := First{
		Ref: &Second{
			Value: 2,
		},
	}
	render, err := util.NewRender("action/doublelink/link.tpl")
	if err != nil {
		panic(err)
	}
	result ,err := render.Rend(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("render result is \n%s\n", result)

	f.Ref.Value = 3
	result ,err = render.Rend(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("render result is \n%s\n", result)
}
