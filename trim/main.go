package main

import (
	"fmt"

	"github.com/leeshun/template-learning/util"
)

type pair struct {
	First int
	Second int
}

func main() {
	p := &pair{
		First:  10,
		Second: 2,
	}
	result, err := util.Execute("trim/trim.tpl", p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("render result is \n%s\n", result)
}
