package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3}
	want := []int{2, 3, 1}

	result := rotate_ints_1(s)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got %v, want %v", result, want)
	}
}

func TestRotate2(t *testing.T) {
	s := []int{1, 2, 3}
	want := []int{3, 1, 2}

	result := rotate_ints_2(s)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got %v, want %v", result, want)
	}	
}
