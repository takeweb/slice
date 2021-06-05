package slice_util

import (
	"fmt"
)

// 最後に要素を追加
func Push(a *interface{}, values ...interface{}) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		for _, value := range values {
			if value, ok := value.(int); ok {
				array = PushInt(array, value)
			}
		}
		i = interface{}(array)
		*a = i
	case []string:
		for _, value := range values {
			if value, ok := value.(string); ok {
				array = PushString(array, value)
			}
		}
		i = interface{}(array)
		*a = i
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
	// return rtn
}

// 最後の要素を削除
func Pop(a *interface{}) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		array = PopInt(array)
		i = interface{}(array)
		*a = i
	case []string:
		array = PopString(array)
		i = interface{}(array)
		*a = i
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 最初に要素を追加
func Unshift(a *interface{}, v interface{}) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		if value, ok := v.(int); ok {
			array = UnshiftInt(array, value)
			i = interface{}(array)
			*a = i
		}
	case []string:
		if value, ok := v.(string); ok {
			array = UnshiftString(array, value)
			i = interface{}(array)
			*a = i
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 最初の要素を削除
func Shift(a *interface{}) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		array = ShiftInt(array)
		i = interface{}(array)
		*a = i
	case []string:
		array = ShiftString(array)
		i = interface{}(array)
		*a = i
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 指定位置に追加
func Insert(a *interface{}, p int, v interface{}) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		if value, ok := v.(int); ok {
			array = InsertInt(array, p, value)
			i = interface{}(array)
			*a = i
		}
	case []string:
		if value, ok := v.(string); ok {
			array = InsertString(array, p, value)
			i = interface{}(array)
			*a = i
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 指定位置を削除
func Remove(a *interface{}, p int) {
	var i interface{}
	switch array := (*a).(type) {
	case []int:
		array = RemoveInt(array, p)
		i = interface{}(array)
		*a = i
	case []string:
		array = RemoveString(array, p)
		i = interface{}(array)
		*a = i
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}
