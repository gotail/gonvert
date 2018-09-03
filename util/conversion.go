package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// type conversion
func convertor(value string, structType string) (reflect.Value, error) {
	fmt.Println(value, structType)
	switch structType {
	case "[]int": //TODO 难道只能这么实现？？？需要思考更好的办法
		value = strings.Trim(value, "[]")
		valueSlice := strings.Split(value, " ")

		var returnSlice []int
		for _, k := range valueSlice {
			i, _ := strconv.Atoi(k)
			returnSlice = append(returnSlice, i)
		}
		return reflect.ValueOf(returnSlice), nil
	case "int":
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	case "uint8":
		i, err := strconv.ParseUint(value, 10, 64)
		return reflect.ValueOf(uint8(i)), err
	case "int8":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	case "int32":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	case "int64":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	case "string":
		return reflect.ValueOf(value), nil
	case "time.Time":
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	case "Time":
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	case "float32":
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	case "float64":
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	return reflect.ValueOf(value), errors.New("unknow：" + structType)
}
