package slice

import (
	"reflect"
	"testing"
)

// 最後に要素を追加のテスト
func TestPushString(t *testing.T) {
	a := []string{"10", "20", "30", "40", "50"}
	a = PushString(a, "1000")
	b := []string{"10", "20", "30", "40", "50", "1000"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最後の要素を削除のテスト
func TestPopString(t *testing.T) {
	a := []string{"10", "20", "30", "40", "50", "1000"}
	a = PopString(a)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最初に要素を追加のテスト
func TestUnshiftString(t *testing.T) {
	a := []string{"10", "20", "30", "40", "50"}
	a = UnshiftString(a, "1000")
	b := []string{"1000", "10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 最初の要素を削除のテスト
func TestShiftString(t *testing.T) {
	a := []string{"1000", "10", "20", "30", "40", "50"}
	a = ShiftString(a)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 指定位置に要素を追加のテスト
func TestInsertString(t *testing.T) {
	a := []string{"10", "20", "30", "40", "50"}
	a = InsertString(a, 2, "1000")
	b := []string{"10", "20", "1000", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}

// 指定位置の要素を削除のテスト
func TestRemoveString(t *testing.T) {
	a := []string{"10", "20", "1000", "30", "40", "50"}
	a = RemoveString(a, 2)
	b := []string{"10", "20", "30", "40", "50"}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s and %s is not equal", a, b)
	}
}
