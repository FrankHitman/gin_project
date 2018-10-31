package utils

import (
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[SnakeString(t.Field(i).Name)] = v.Field(i).Interface()
	}
	//fmt.Println(data)
	return data
}