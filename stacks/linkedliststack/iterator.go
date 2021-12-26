package linkedliststack

import "gods/containers"

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

type Iterator struct {
	stack *Stack
	index int
}

func (stack *Stack) Iterator() Iterator {
	return Iterator{stack: stack, index: -1}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.stack.list.Get(iterator.index) // in reverse (LIFO)
	return value
}

func (iterator *Iterator) Index() int {
	return iterator.index
}

func (iterator *Iterator) Begin() {
	iterator.index = -1
}

func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}
