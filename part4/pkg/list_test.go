package list

import (
	"testing"
)

func TestList_Push(t *testing.T) {
	l := New()
	l.Push(Elem{Val: 3})
	l.Push(Elem{Val: 2})
	l.Push(Elem{Val: 1})

	got := l.String()
	want := "1 2 3"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestList_Pop(t *testing.T) {
	l := New()
	l.Push(Elem{Val: 4})
	l.Push(Elem{Val: 3})
	l.Push(Elem{Val: 2})
	l.Push(Elem{Val: 1})

	got := l.Pop().String()
	want := "2 3 4"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}

	got = l.Pop().Pop().Pop().String()
	want = ""
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestList_Delete(t *testing.T) {
	l := New()
	l.Push(Elem{Val: 4})
	l.Push(Elem{Val: 3})
	l.Push(Elem{Val: 2})
	l.Push(Elem{Val: 1})

	l.Delete()
	got := l.String()
	want := "1 2 3"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}

	l.Delete()
	l.Delete()
	l.Delete()
	got = l.String()
	want = ""
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestList_Reverse(t *testing.T) {
	l := New()

	got := l.Reverse().String()
	want := ""
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}

	l.Push(Elem{Val: 1})

	got = l.Reverse().String()
	want = "1"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}

	l.Push(Elem{Val: 2})
	l.Push(Elem{Val: 3})

	got = l.Reverse().String()
	want = "1 2 3"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
