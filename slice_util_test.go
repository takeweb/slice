package slice_util

import (
	"reflect"
	"testing"
)

// 最後に要素(string)を追加のテスト
func TestPushS(t *testing.T) {
	var a interface{} = []string{"10", "20", "30", "40", "50"}
	Push(&a, "1000", "2000")
	b := []string{"10", "20", "30", "40", "50", "1000", "2000"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最後に要素(int)追加のテスト
func TestPushI(t *testing.T) {
	var a interface{} = []int{10, 20, 30, 40, 50}
	Push(&a, 1000)
	b := []int{10, 20, 30, 40, 50, 1000}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最後の要素(string)を削除のテスト
func TestPopS(t *testing.T) {
	var a interface{} = []string{"10", "20", "30", "40", "50", "1000"}
	Pop(&a)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最後の要素(int)を削除のテスト
func TestPopI(t *testing.T) {
	var a interface{} = []int{10, 20, 30, 40, 50, 1000}
	Pop(&a)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最初に要素(string)を追加のテスト
func TestUnshiftS(t *testing.T) {
	var a interface{} = []string{"10", "20", "30", "40", "50"}
	Unshift(&a, "1000")
	b := []string{"1000", "10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最初に要素(int)を追加のテスト
func TestUnshiftI(t *testing.T) {
	var a interface{} = []int{10, 20, 30, 40, 50}
	Unshift(&a, 1000)
	b := []int{1000, 10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最初の要素(string)を削除のテスト
func TestShiftS(t *testing.T) {
	var a interface{} = []string{"1000", "10", "20", "30", "40", "50"}
	Shift(&a)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最初の要素(int)を削除のテスト
func TestShiftI(t *testing.T) {
	var a interface{} = []int{1000, 10, 20, 30, 40, 50}
	Shift(&a)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 指定位置に要素(string)を追加のテスト
func TestInsertS(t *testing.T) {
	var a interface{} = []string{"10", "20", "30", "40", "50"}
	Insert(&a, 2, "1000")
	b := []string{"10", "20", "1000", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 指定位置に要素(int)を追加のテスト
func TestInsertI(t *testing.T) {
	var a interface{} = []int{10, 20, 30, 40, 50}
	Insert(&a, 2, 1000)
	b := []int{10, 20, 1000, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 指定位置の要素(string)を削除のテスト
func TestRemoveS(t *testing.T) {
	var a interface{} = []string{"10", "20", "1000", "30", "40", "50"}
	Remove(&a, 2)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 指定位置の要素を削除のテスト
func TestRemoveI(t *testing.T) {
	var a interface{} = []int{10, 20, 1000, 30, 40, 50}
	Remove(&a, 2)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}
