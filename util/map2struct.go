package util

import (
	"fmt"
	"reflect"
)

// Map2Struct 将Map转换为Struct
func Map2Struct(mapData map[string]interface{}, structInterface interface{}, tagName string) err {
	getType := reflect.TypeOf(structInterface).Elem()
	getValue := reflect.ValueOf(structInterface).Elem()

	for i := 0; i < getValue.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)

		mapKey := field.Tag.Get(tagName)
		v, ok := mapData[mapKey]
		if ok == false {
			continue
		}

		mapValue := reflect.ValueOf(v)
		// 若map中的类型与结构体类型不一致，则需要执行类型转换
		if reflect.TypeOf(mapValue) != value.Type() {
			var err error
			mapValue, err = convertor(fmt.Sprintf("%v", mapValue), value.Type().Name())
			if err != nil {
				return err
			}
		}
		value.Set(mapValue)
	}
	
	return nil
}
