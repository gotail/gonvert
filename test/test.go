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
}

var data = map[string]interface{}{
	"username": "golang",
	"age":      30,
	"company":  "Github",
	// "jh":true,
	"x": 32,
}

func main() {
	person := &Person{}
	gonvert.Map2Struct(data, person)
	fmt.Println(person)

	a, _ := gonvert.Struct2Map(*person)
	fmt.Println(a)
}
