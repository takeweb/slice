package slice_util

import (
	"fmt"
)

// 最後に要素を追加
func Push(a interface{}, values ...interface{}) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		for _, value := range values {
			if value, ok := value.(int); ok {
				array = PushInt(array, value)
			}
		}
		rtn = array
	case []string:
		for _, value := range values {
			if value, ok := value.(string); ok {
				array = PushString(array, value)
			}
		}
		rtn = array
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}

// 最後の要素を削除
func Pop(a interface{}) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		rtn = PopInt(array)
	case []string:
		rtn = PopString(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}

// 最初に要素を追加
func Unshift(a interface{}, v interface{}) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		if value, ok := v.(int); ok {
			rtn = UnshiftInt(array, value)
		}
	case []string:
		if value, ok := v.(string); ok {
			rtn = UnshiftString(array, value)
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}

// 最初の要素を削除
func Shift(a interface{}) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		rtn = ShiftInt(array)
	case []string:
		rtn = ShiftString(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}

// 指定位置に追加
func Insert(a interface{}, p int, v interface{}) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		if value, ok := v.(int); ok {
			array = InsertInt(array, p, value)
		}
		rtn = array
	case []string:
		if value, ok := v.(string); ok {
			array = InsertString(array, p, value)
		}
		rtn = array
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}

// 指定位置を削除
func Remove(a interface{}, p int) interface{} {
	var rtn interface{}
	switch array := a.(type) {
	case []int:
		rtn = RemoveInt(array, p)
	case []string:
		rtn = RemoveString(array, p)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	return rtn
}
