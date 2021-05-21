package slice_util

import (
	"reflect"
	"testing"
)

// 最後に追加のテスト
func TestPushInt(t *testing.T) {
	a := []int{10, 20, 30, 40, 50}
	a = PushInt(a, 1000)
	b := []int{10, 20, 30, 40, 50, 1000}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最後を削除のテスト
func TestPopInt(t *testing.T) {
	a := []int{10, 20, 30, 40, 50, 1000}
	a = PopInt(a)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最初に追加のテスト
func TestUnshiftInt(t *testing.T) {
	a := []int{10, 20, 30, 40, 50}
	a = UnshiftInt(a, 1000)
	b := []int{1000, 10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 最初を削除のテスト
func TestShiftInt(t *testing.T) {
	a := []int{1000, 10, 20, 30, 40, 50}
	a = ShiftInt(a)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 指定位置に追加のテスト
func TestInsertInt(t *testing.T) {
	a := []int{10, 20, 30, 40, 50}
	a = InsertInt(a, 2, 1000)
	b := []int{10, 20, 1000, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}

// 指定位置を削除のテスト
func TestRemoveInt(t *testing.T) {
	a := []int{10, 20, 1000, 30, 40, 50}
	a = RemoveInt(a, 2)
	b := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%d and %d is not equal", a, b)
	}
}
