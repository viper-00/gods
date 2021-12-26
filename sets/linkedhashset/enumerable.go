package linkedhashset

import "gods/containers"

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithIndex = (*Set)(nil)
}

func (set *Set) Each(f func(index int, value interface{})) {
	iterator := set.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

func (set *Set) Map(f func(index int, value interface{}) interface{}) *Set {
	newSet := New()
	iterator := set.Iterator()
	for iterator.Next() {
		newSet.Add(f(iterator.Index(), iterator.Value()))
	}
	return newSet
}

func (set *Set) Select(f func(index int, value interface{}) bool) *Set {
	newSet := New()
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newSet.Add(iterator.Value())
		}
	}
	return newSet
}

func (set *Set) Any(f func(index int, value interface{}) bool) bool {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

func (set *Set) All(f func(index int, value interface{}) bool) bool {
	iterator := set.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

func (set *Set) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}
