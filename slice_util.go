package slice_util

import (
	"fmt"
)

// 最後に要素を追加
func Push(a *interface{}, values ...interface{}) {
	switch array := (*a).(type) {
	case []int:
		for _, value := range values {
			if value, ok := value.(int); ok {
				array = PushInt(array, value)
			}
		}
		*a = interface{}(array)
	case []string:
		for _, value := range values {
			if value, ok := value.(string); ok {
				array = PushString(array, value)
			}
		}
		*a = interface{}(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 最後の要素を削除
func Pop(a *interface{}) {
	switch array := (*a).(type) {
	case []int:
		array = PopInt(array)
		*a = interface{}(array)
	case []string:
		array = PopString(array)
		*a = interface{}(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 最初に要素を追加
func Unshift(a *interface{}, v interface{}) {
	switch array := (*a).(type) {
	case []int:
		if value, ok := v.(int); ok {
			array = UnshiftInt(array, value)
			*a = interface{}(array)
		}
	case []string:
		if value, ok := v.(string); ok {
			array = UnshiftString(array, value)
			*a = interface{}(array)
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 最初の要素を削除
func Shift(a *interface{}) {
	switch array := (*a).(type) {
	case []int:
		array = ShiftInt(array)
		*a = interface{}(array)
	case []string:
		array = ShiftString(array)
		*a = interface{}(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 指定位置に追加
func Insert(a *interface{}, p int, v interface{}) {
	switch array := (*a).(type) {
	case []int:
		if value, ok := v.(int); ok {
			array = InsertInt(array, p, value)
			*a = interface{}(array)
		}
	case []string:
		if value, ok := v.(string); ok {
			array = InsertString(array, p, value)
			*a = interface{}(array)
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}

// 指定位置を削除
func Remove(a *interface{}, p int) {
	switch array := (*a).(type) {
	case []int:
		array = RemoveInt(array, p)
		*a = interface{}(array)
	case []string:
		array = RemoveString(array, p)
		*a = interface{}(array)
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", a)
	}
}
