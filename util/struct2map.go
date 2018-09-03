package util

import "reflect"

// Struct2Map 将Struct转换为Map
func Struct2Map(structInterface interface{}, tagName string) map[string]interface{} {
	getType := reflect.TypeOf(structInterface).Elem()
	getValue := reflect.ValueOf(structInterface).Elem()

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
