package slice_util

import "reflect"

// General is all type data.
type General_s []interface{}

// 最後に追加
func Push(a []General_s, v ...General) []General {
	var dd General
	if reflect.TypeOf(a) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = a.([]string)
	}
	if reflect.TypeOf(a).Kind() == reflect.String {
		b := a.([]string)
		w := v.(string)
		return PushString(b, w)
	}
	return dd
}

// 最後を削除
func Pop(a []string) []string {
	return a[:len(a)-1]
}

// 最初に追加
func Unshift(a []string, v string) []string {
	return append([]string{v}, a...)
}

// 最初を削除
func Shift(a []string) []string {
	return a[1:]
}

// 指定位置に追加
func Insert(a []string, p int, v string) []string {
	a = append(a, "0")
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

// 指定位置を削除
func Remove(a []string, p int) []string {
	return append(a[:p], a[p+1:]...)
}
