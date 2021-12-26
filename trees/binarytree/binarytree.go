package binarytree

import (
	"fmt"
	"gods/trees"
	"gods/utils"
)

func assertTreeImplementation() {
	var _ trees.Tree = (*Tree)(nil)
}

type Tree struct {
	Root       *Node
	size       int
	Comparator utils.Comparator
}

type Node struct {
	Key    interface{}
	Value  interface{}
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

func (tree *Tree) Put(key interface{}, value interface{}) {
	var insertedNode *Node
	if tree.Root == nil {
		tree.Root = &Node{Key: key, Value: value}
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.Comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node{Key: key, Value: value}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node{Key: key, Value: value}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.size++
}

func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

func (tree *Tree) Remove(key interface{}) {
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		var child *Node
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		tree.replaceNode(node, child)
	}
	tree.size--
}

func (tree *Tree) Empty() bool {
	return tree.size == 0
}

func (tree *Tree) Size() int {
	return tree.size
}

func (tree *Tree) Keys() []interface{} {
	keys := make([]interface{}, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

func (tree *Tree) Values() []interface{} {
	values := make([]interface{}, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

func (tree *Tree) Left() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

func (tree *Tree) Right() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

func (tree *Tree) Clear() {
	tree.Root = nil
	tree.size = 0
}

func (tree *Tree) String() string {
	str := "BinaryTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "|   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (tree *Tree) lookup(key interface{}) *Node {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare > 0:
			node = node.Right
		case compare < 0:
			node = node.Left
		}
	}
	return nil
}

func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}
