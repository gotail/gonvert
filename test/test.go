package main

import (
	"fmt"

	"github.com/gotail/gonvert"
)

type Person struct {
	Username string `gonvert:"username"`
	Age      int    `gonvert:"age"`
	Company  string `gonvert:"company"`
	Jh       bool   `gonvert:"jh"`
	X        int64  `gonvert:"x"`
	Y        byte   `gonvert:"y"`
	Z        []int  `gonvert:"z"`
}

var data = map[string]interface{}{
	"username": "golang",
	// "age":      30,
	// "company":  "Github",
	"jh":true,
	"x": 0,
	// "y": 'a',
	//"z": []int{1, 2, 3},
}

func main() {
	person := Person{}
	err := gonvert.Map2Struct(data, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
	//
	//a, _ := gonvert.Struct2Map(&person)
	//fmt.Println(a)
}
