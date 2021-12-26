package sets

import "gods/containers"

// Package sets provides an abstract Set intterface.
//
// In computer science, a set is an abstract data type tha can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set.
// Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.
//

// Set interface that all sets implement.
type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool

	containers.Container
}
