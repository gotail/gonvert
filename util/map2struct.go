package util

import (
	"fmt"
	"reflect"
)

// Map2Struct 将Map转换为Struct
func Map2Struct(data map[string]interface{}, obj interface{}, tagName string) {

	t := reflect.ValueOf(obj).Elem()
	tType := t.Type()

	for i := 0; i < t.NumField(); i++ {
		mapKey := tType.Field(i).Tag.Get(tagName)
		v, ok := data[mapKey]
		value := reflect.ValueOf(v)
		if ok {
			// 若map中的类型与结构体类型不一致，则需要执行类型转换
			if reflect.TypeOf(value) != t.Field(i).Type() {
				var err error
				value, err = convertor(fmt.Sprintf("%v", value), t.Field(i).Type().Name())
				if err != nil {
					panic(err)
				}
			}
			t.Field(i).Set(value)
		}

	}
}
