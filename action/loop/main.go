package main

import (
	"fmt"

	"github.com/leeshun/template-learning/util"
)

func main() {
	r, err := util.NewRender("action/loop/range.tpl")
	if err != nil {
		panic(err)
	}

	fmt.Printf("empty slice\n")
	result, err := r.Rend(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	val := struct {
		Value []int
	}{
		Value: []int{1,2,3},
	}
	fmt.Println("integer slice")
	result, err = r.Rend(val)
	if err!= nil {
		panic(err)
	}
	fmt.Println(result)

	mVal := struct {
		Value map[int]int
	}{
		Value: map[int]int{
			1:10,
			2:20,
			3:30,
		},
	}
	fmt.Println("integer map")
	result, err = r.Rend(mVal)
	if err!= nil {
		panic(err)
	}
	fmt.Println(result)

	chVal := struct {
		Value chan int
	}{
		Value: make(chan int, 10),
	}

	for i := 0; i < 10; i++ {
		chVal.Value <- i*i
	}
	close(chVal.Value)

	fmt.Println("integer channel")
	result, err = r.Rend(chVal)
	if err!= nil {
		panic(err)
	}
	fmt.Println(result)
}
