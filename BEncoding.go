package main

import (
	"fmt"
	"reflect"
)

func Encode(obj interface{}) string {
	return encode(reflect.ValueOf(obj))
}

func encode(value reflect.Value) string {
	var result string
	kind := value.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result += fmt.Sprintf("i%de", value.Int())
	case reflect.String:
		result += fmt.Sprintf("%d:%s", value.Len(), value.String())
	case reflect.Slice:
		if value.Len() > 0 {
			result += "l"
			for i := 0; i < value.Len(); i++ {
				result += encode(value.Index(i)) //这里写value.Slice(i,i)就不行
			}
			result += "e"
		}
	case reflect.Map:
		if value.Len() > 0 {
			result += "d"
			if value.Type().Key().Kind() == reflect.String {
				for _, key := range value.MapKeys() {
					result += encode(key)
					if value.MapIndex(key).Kind() == reflect.Interface {
						//fmt.Println("...1...")
						result += encode(value.MapIndex(key).Elem())
					} else {
						//fmt.Println("...2...")
						result += encode(value.MapIndex(key))
					}
				}
			} else {
				fmt.Printf("key of map is not string\n")
				return ""
			}
			result += "e"
		}
	}
	return result
}

func main() {
	/*
		obj := map[string]interface{}{
			"o1": []string{"str1", "str2", "str3"},
			"o2": []string{"str4", "str5"},
			"o3": []int{12, 32},
		}
	*/

	obj := map[string]interface{}{
		"name": "create chen",
		"age":  23,
	}

	var i int = 30
	var str string = "hello"
	var ints []int = []int{12, 32, 1, 45, 2, 8}
	var strs []string = []string{"str1", "str2", "str3"}

	fmt.Println("---encode int---")
	fmt.Println(Encode(i))
	fmt.Println("---encode string---")
	fmt.Println(Encode(str))
	fmt.Println("---encode int list---")
	fmt.Println(Encode(ints))
	fmt.Println("---encode string list---")
	fmt.Println(Encode(strs))
	fmt.Println("---encode map[string]object---")
	fmt.Println(Encode(obj))
}
