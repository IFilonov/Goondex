package sort_test

import (
	"sort"
	"testing"
)

var ints = [...]int{4, 6, 7, 3, 1}

// var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
// var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestSortIntSlice(t *testing.T) {
	sort.Ints(ints[:])
	expected := [...]int{1, 3, 4, 6, 7}
	if ints != expected {
		t.Errorf("expected '%v' but got '%v'", expected, ints)
	}
	// for i, _ := range ints {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	if ints[i-1] > ints[i] {
	// 		t.Errorf("expected '%v' but got '%v'", ints, ints)
	// 		return
	// 	}
	// }
}
