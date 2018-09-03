package gonvert

import (
	"github.com/gotail/gonvert/util"
)

const tagName string = "gonvert"

// Struct2Map 将Struct转换为Map
func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	return util.Struct2Map(obj, tagName), nil
}

// Map2Struct 将Map转换为Struct
func Map2Struct(data map[string]interface{}, obj interface{}) error {
	util.Map2Struct(data, obj, tagName)
	return nil
}
