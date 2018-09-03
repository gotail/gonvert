刚学Go，练练手，万一能做好呢？希望集成各种转换功能。

目前该项目都属于学习阶段，各位不要应用到自己的项目中，如果有一天我认为成熟了可以用于开发了，我会通知。

# gonvert

Provide the transformation of map and struct for the go language

## 安装
```go
go get -u -t github.com/gotail/gonvert
```

## Struct 转换为 Map

```go
package main

import (
	"fmt"

	"github.com/gotail/gonvert"
)

type Person struct {
	Username string `gonvert:"username"`
	Age      int    `gonvert:"age"`
	Company  string `gonvert:"company"`
}

func main() {
	person := Person{Username: "golang", Age: 10, Company: "Github"}
	PersonMap := gonvert.Struct2Map(&person)
	fmt.Println(PersonMap)
}

```
## Map转换为 Struct
```go
package main

import (
	"fmt"

	"github.com/gotail/gonvert"
)

type Person struct {
	Username string `gonvert:"username"`
	Age      int `gonvert:"age"`
	Company  string `gonvert:"company"`
	Jh bool	`gonvert:"jh"`
}

var data = map[string]interface{}{
	"username": "golang",
	"age":      30,
	"company":  "Github",
	"jh":true,
}

func main() {
	person := &Person{}
	gonvert.Map2Struct(data, person)
	fmt.Println(person)
}

```
