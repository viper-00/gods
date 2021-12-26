package stacks

import "gods/containers"

// Package stacks provides an abstract Stack interface.
//
// In computer science, a stack is an abstract data type that serves as a collection of element, with two principal operations: push, whcih adds an element to the collection, and pop, which removes the most recently added element that was not yet removed.
// The order in which elements come off a stack gives rise to its alternative name, LIFO (for last in, first out). Additionally, a peek operation may give access to the top without modifying the stack.

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	containers.Container
}
