// Package stringer provides exercises for implementing the Stringer interface,
// enabling custom string formatting.
package stringer

/*
Task: Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and
many others) look for this interface to print values.

Implement the String() method for the Student struct.

A struct

Student{ID: 42, FirstName: John, LastName: Doe, Age: 25}

should be printed as

"Student ID: 42. Name: Doe, John. Age: 25.
*/

// Student holds information about a student.
type Student struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func (s Student) String() string {
	return ""
}
