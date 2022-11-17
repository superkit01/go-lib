package slice

import (
	"testing"
)

func TestMinValue(t *testing.T) {
	expected := 1
	v, err := MinValue([]int{1, 2, 3, 4, 7, 5})
	if err != nil || v != expected {
		t.Errorf("result:%v; expected:%v ", v, expected)
	}
}

func TestMaxValue(t *testing.T) {
	expected := 7
	v, err := MaxValue([]int{1, 2, 3, 4, 7, 5})
	if err != nil || v != expected {
		t.Errorf("result:%v; expected:%v ", v, expected)
	}
}
