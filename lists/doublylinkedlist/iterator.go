package doublylinkedlist

import "gods/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

type Iterator struct {
	list    *List
	index   int
	element *element
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index != 0 {
		iterator.element = iterator.element.next
	} else {
		iterator.element = iterator.list.first
	}
	return true
}

func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == iterator.list.size-1 {
		iterator.element = iterator.list.last
	} else {
		iterator.element = iterator.element.prev
	}
	return true
}

func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

func (iterator *Iterator) Index() int {
	return iterator.index
}

func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
