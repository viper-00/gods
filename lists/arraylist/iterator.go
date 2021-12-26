package arraylist

import "gods/containers"

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

// Iterator holding the iterator's state
type Iterator struct {
	list  *List
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

func (Iterator *Iterator) Index() int {
	return Iterator.index
}

func (iterator *Iterator) Begin() {
	iterator.index = -1
}

func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
}

func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
