package person

import (
	"io"
)

type Employee struct {
	years int
}

type Customer struct {
	years int
}

func (e *Employee) age() int {
	return e.years
}

func (c *Customer) age() int {
	return c.years
}

type person interface {
	age() int
}

// 1 -------------------------------------

func MaxAge(persons ...person) int {
	var maxAge int
	for _, person := range persons {
		if maxAge < person.age() {
			maxAge = person.age()
		}
	}
	return maxAge
}

// 2 -------------------------------------

func OldestPerson(persons ...any) any {
	var maxAge, personAge int
	var oldestPerson any
	for _, person := range persons {
		switch p := person.(type) {
		case Employee:
		case Customer:
			personAge = p.years
		default:
			personAge = 0
		}
		if maxAge < personAge {
			maxAge = personAge
			oldestPerson = person
		}
	}
	return oldestPerson
}

// 3 -------------------------------------

func StringWriter(w io.Writer, vars ...any) error {
	for _, anyVar := range vars {
		stringVar, ok := anyVar.(string)
		if ok {
			_, err := w.Write([]byte(stringVar))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
