package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	list := ListElement[string]{nil, "first"}
	secondElement := ListElement[string]{nil, "second"}
	list.next = &secondElement
	thirdElement := ListElement[string]{nil, "third elem"}
	secondElement.next = &thirdElement

	fmt.Println(list.lastElement().val)
}

type ListElement[T any] struct {
	next *ListElement[T]
	val  T
}

func (list *ListElement[T]) lastElement() *ListElement[T] {
	var p = list
	for p.next != nil {
		p = p.next
	}
	return p
}
