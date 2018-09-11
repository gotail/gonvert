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

		if value.IsValid() == false {
			continue
		}

		mapKey := field.Tag.Get(tagName)
		v, ok := mapData[mapKey]
		if ok == false || v == nil {
			continue
		}

		mapValue := reflect.ValueOf(v)
		structType := value.Type().Name()
		// 若map中的类型与结构体类型不一致，则需要执行类型转换
		if reflect.TypeOf(mapValue) != value.Type() && structType != "" {
			var err error
			mapValue, err = convertor(fmt.Sprintf("%v", mapValue), structType)
			if err != nil {
				//fmt.Printf("发生了错误：%v\n", err)
				return err
				continue
			}
		}
		value.Set(mapValue)
	}

	return nil
}
