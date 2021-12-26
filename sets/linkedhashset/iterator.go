package linkedhashset

import (
	"gods/containers"
	"gods/lists/doublylinkedlist"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

type Iterator struct {
	iterator doublylinkedlist.Iterator
}

func (set *Set) Iterator() Iterator {
	return Iterator{iterator: set.ordering.Iterator()}
}

func (iterator *Iterator) Next() bool {
	return iterator.iterator.Next()
}

func (iterator *Iterator) Prev() bool {
	return iterator.iterator.Prev()
}

func (iterator *Iterator) Value() interface{} {
	return iterator.iterator.Value()
}

func (iterator *Iterator) Index() int {
	return iterator.iterator.Index()
}

func (iterator *Iterator) Begin() {
	iterator.iterator.Begin()
}

func (iterator *Iterator) End() {
	iterator.iterator.End()
}

func (iterator *Iterator) First() bool {
	return iterator.iterator.First()
}

func (iterator *Iterator) Last() bool {
	return iterator.iterator.Last()
}
