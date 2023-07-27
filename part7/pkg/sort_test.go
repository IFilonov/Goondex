package sort_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

var ints = [...]int{4, 6, 7, 3, 1}

const MAX_RAND_LIMIT = 1000

func TestSortIntSlice(t *testing.T) {
	sort.Ints(ints[:])
	expected := [...]int{1, 3, 4, 6, 7}
	if ints != expected {
		t.Errorf("expected '%v' but got '%v'", expected, ints)
	}
}

func TestSortStringSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{
			name: "Тест 1",
			s:    []string{"C", "B", "A"},
			want: []string{"A", "B", "C"},
		},
		{
			name: "Тест 2",
			s:    []string{"Z", "X", "C"},
			want: []string{"C", "X", "Z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.s)
			if got := tt.s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected '%v' but got '%v'", got, tt.want)
			}
		})
	}
}

func generateRandomIntArray() []int {
	var a []int
	for i := 0; i < 10000; i++ {
		a = append(a, rand.Intn(MAX_RAND_LIMIT))
	}
	return a
}

func generateRandomFloatArray() []float64 {
	var a []float64
	for i := 0; i < 10000; i++ {
		a = append(a, rand.Float64())
	}
	return a
}

func BenchmarkIntSort(b *testing.B) {
	a := generateRandomIntArray()
	for n := 0; n < b.N; n++ {
		sort.Ints(a)
	}
}

func BenchmarkFloat64Sort(b *testing.B) {
	a := generateRandomFloatArray()
	for n := 0; n < b.N; n++ {
		sort.Float64s(a)
	}
}
