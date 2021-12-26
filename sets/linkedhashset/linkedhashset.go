package linkedhashset

import (
	"fmt"
	"gods/lists/doublylinkedlist"
	"gods/sets"
	"strings"
)

// Package linkedhashset is a set that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

type Set struct {
	table    map[interface{}]struct{}
	ordering *doublylinkedlist.List
}

var itemExists = struct{}{}

func New(values ...interface{}) *Set {
	set := &Set{
		table:    make(map[interface{}]struct{}),
		ordering: doublylinkedlist.New(),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		if _, ok := set.table[item]; !ok {
			set.table[item] = itemExists
			set.ordering.Append(item)
		}
	}
}

func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		if _, ok := set.table[item]; ok {
			delete(set.table, item)
			index := set.ordering.IndexOf(item)
			set.ordering.Remove(index)
		}
	}
}

func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := set.table[item]; !ok {
			return false
		}
	}
	return true
}

func (set *Set) Empty() bool {
	return set.Size() == 0
}

func (set *Set) Size() int {
	// len(set.table) or set.ordering.Size()
	return set.ordering.Size()
}

func (set *Set) Clear() {
	set.table = make(map[interface{}]struct{})
	set.ordering.Clear()
}

func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	it := set.Iterator()
	for it.Next() {
		values[it.Index()] = it.Value()
	}
	return values
}

func (set *Set) String() string {
	str := "LinkedHashSet\n"
	items := []string{}
	it := set.Iterator()
	for it.Next() {
		items = append(items, fmt.Sprintf("%v", it.Value()))
	}
	str += strings.Join(items, ", ")
	return str
}
