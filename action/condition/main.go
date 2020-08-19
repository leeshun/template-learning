package main

import (
	"fmt"
	"os"
	"text/template"
)

type singleValue struct {
	Value *int
}

func singleCondition() {
	tpl, err := template.ParseFiles("action/condition/only_if.tpl")
	if err != nil {
		panic(err)
	}
	val := 2
	s := &singleValue{
		Value:&val,
	}
	fmt.Println("single case with positive value")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}

	val = -1
	fmt.Println("single case with negative value")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}

	s.Value = nil
	fmt.Println("single value with empty pointer")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}

func doubleCondition() {
	tpl, err := template.ParseFiles("action/condition/if_else.tpl")
	if err != nil {
		panic(err)
	}
	val := 10
	s := &singleValue{
		Value: &val,
	}
	fmt.Println("double case with positive value")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}

	val = 0
	fmt.Println("double case with zero")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}

	s.Value = nil
	fmt.Println("double case with empty pointer")
	err = tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}

func main() {
	singleCondition()

	doubleCondition()
}
