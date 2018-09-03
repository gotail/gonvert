package util

import "reflect"

// Struct2Map 将Struct转换为Map
func Struct2Map(structInterface interface{}, tagName string) map[string]interface{} {
	getType := reflect.TypeOf(structInterface)
	getValue := reflect.ValueOf(structInterface)

	if getType.Kind() == reflect.Ptr { // 判断其是否是指针
		getType = getType.Elem()
		getValue = getValue.Elem()
	}

	mapData := make(map[string]interface{})

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()

		keyName := field.Tag.Get(tagName)
		if keyName == "" {
			keyName = field.Name
		}
		mapData[keyName] = value
	}

	return mapData
}
