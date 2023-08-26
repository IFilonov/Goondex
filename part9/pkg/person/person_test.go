package person

import (
	"bytes"
	"testing"
)

// 1 -------------------------------------

func Test_MaxAge(t *testing.T) {
	test_data := []person{
		&Customer{years: 1},
		&Employee{years: 2},
		&Customer{years: 5},
		&Employee{years: 4},
		&Customer{years: 3},
	}

	want := 5
	got := MaxAge(test_data...)
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

// 2 -------------------------------------

func Test_OldestPerson(t *testing.T) {
	test_data := []any{
		Customer{years: 1},
		Employee{years: 2},
		Customer{years: 5},
		Employee{years: 4},
		Customer{years: 3},
	}

	want := Customer{years: 5}
	got := OldestPerson(test_data...)
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

// 3 -------------------------------------

func Test_StringWriter(t *testing.T) {
	test_data := []any{
		Employee{years: 2},
		"24",
		54,
		"46",
		Customer{years: 3},
	}

	var b bytes.Buffer
	err := StringWriter(&b, test_data...)
	if err != nil {
		t.Fatalf("ошибка выполнения StringWriter: %s", err.Error())
	}
	want := []byte("2446")
	got := b.Bytes()
	if !bytes.Equal(got, want) {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
