package binarytree

import "gods/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)
}

type Iterator struct {
	tree     *Tree
	node     *Node
	position position
}

const (
	begin, between, end position = 0, 1, 2
)

type position byte

func (tree *Tree) Iterator() Iterator {
	return Iterator{tree: tree, node: nil, position: begin}
}

func (iterator *Iterator) Next() bool {
	if iterator.position == end {
		goto end
	}
	if iterator.position == begin {
		left := iterator.tree.Left()
		if left == nil {
			goto end
		}
		iterator.node = left
		goto between
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if iterator.tree.Comparator(node.Key, iterator.node.Key) <= 0 {
				goto between
			}
		}
	}
end:
	iterator.node = nil
	iterator.position = end
	return false
between:
	iterator.position = between
	return true

}

func (iterator *Iterator) Prev() bool {
	if iterator.position == begin {
		goto begin
	}
	if iterator.position == end {
		right := iterator.tree.Right()
		if right == nil {
			goto begin
		}
		iterator.node = right
		goto between
	}
	if iterator.node.Left != nil {
		iterator.node = iterator.node.Left
		for iterator.node.Right != nil {
			iterator.node = iterator.node.Right
		}
		goto between
	}
	if iterator.node.Parent != nil {
		node := iterator.node
		for iterator.node.Parent != nil {
			iterator.node = iterator.node.Parent
			if iterator.tree.Comparator(node.Key, iterator.node.Key) >= 0 {
				goto between
			}
		}
	}

begin:
	iterator.node = nil
	iterator.position = begin
	return false
between:
	iterator.position = between
	return true
}

func (iterator *Iterator) Value() interface{} {
	return iterator.node.Value
}

func (iterator *Iterator) Key() interface{} {
	return iterator.node.Key
}

func (iterator *Iterator) Begin() {
	iterator.node = nil
	iterator.position = begin
}

func (iterator *Iterator) End() {
	iterator.node = nil
	iterator.position = end
}

func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
