package util

import (
	"errors"
	"fmt"
	"reflect"
)

// Map2Struct 将Map转换为Struct
func Map2Struct(mapData map[string]interface{}, structInterface interface{}, tagName string) error {
	getType := reflect.TypeOf(structInterface)
	getValue := reflect.ValueOf(structInterface)

	if getType.Kind() == reflect.Ptr { // 判断其是否是指针
		getType = getType.Elem()
		getValue = getValue.Elem()
	} else {
		return errors.New("请传入指针")
	}

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

			var structType = value.Type().Name()
			if structType == "" {
				structType = value.Type().String()
			}

			mapValue, err = convertor(fmt.Sprintf("%v", mapValue), structType)
			if err != nil {
				return err
			}
		}
		value.Set(mapValue)
	}

	return nil
}
