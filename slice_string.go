package slice

// 最後に追加
func PushString(a []string, v ...string) []string {
	return append(a, v...)
}

// 最後を削除
func PopString(a []string) []string {
	return a[:len(a)-1]
}

// 最初に追加
func UnshiftString(a []string, v string) []string {
	return append([]string{v}, a...)
}

// 最初を削除
func ShiftString(a []string) []string {
	return a[1:]
}

// 指定位置に追加
func InsertString(a []string, p int, v string) []string {
	a = append(a, "0")
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

// 指定位置を削除
func RemoveString(a []string, p int) []string {
	return append(a[:p], a[p+1:]...)
}
