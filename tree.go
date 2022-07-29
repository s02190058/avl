package avl

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	left, right *node[T]
	key         T
	height      int
}

func (p *node[T]) getHeight() int {
	if p == nil {
		return 0
	}
	return p.height
}

func (p *node[T]) balanceFactor() int {
	return p.right.getHeight() - p.left.getHeight()
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func (p *node[T]) fix() {
	p.height = max(p.left.getHeight(), p.right.getHeight()) + 1
}

func rotateLeft[T constraints.Ordered](p *node[T]) *node[T] {
	ret := p.right
	p.right = ret.left
	ret.left = p
	p.fix()
	ret.fix()
	return ret
}

func rotateRight[T constraints.Ordered](p *node[T]) *node[T] {
	ret := p.left
	p.left = ret.right
	ret.right = p
	p.fix()
	ret.fix()
	return ret
}

func balance[T constraints.Ordered](p *node[T]) *node[T] {
	p.fix()
	switch p.balanceFactor() {
	case -2:
		if p.left.balanceFactor() > 0 {
			p.left = rotateLeft(p.left)
		}
		return rotateRight(p)
	case 2:
		if p.right.balanceFactor() < 0 {
			p.right = rotateRight(p.right)
		}
		return rotateLeft(p)
	default:
		return p
	}
}

func insert[T constraints.Ordered](p *node[T], key T) *node[T] {
	if p == nil {
		return &node[T]{
			key: key,
		}
	}
	if key < p.key {
		p.left = insert(p.left, key)
	} else {
		p.right = insert(p.right, key)
	}
	return balance(p)
}

func findMin[T constraints.Ordered](p *node[T]) *node[T] {
	if p.left == nil {
		return p
	}
	return findMin(p.left)
}

func eraseMin[T constraints.Ordered](p *node[T]) *node[T] {
	if p.left == nil {
		return p.right
	}
	p.left = eraseMin(p.left)
	return balance(p)
}

func erase[T constraints.Ordered](p *node[T], key T) *node[T] {
	if p == nil {
		return nil
	}
	if key < p.key {
		p.left = erase(p.left, key)
	} else if key > p.key {
		p.right = erase(p.right, key)
	} else {
		l, r := p.left, p.right
		if r == nil {
			return l
		}
		p = findMin(r)
		p.right = eraseMin(r)
		p.left = l
	}
	return balance(p)
}

type Tree[T constraints.Ordered] struct {
	root *node[T]
}

func (t Tree[T]) Height() int {
	return t.root.getHeight()
}

func (p *node[T]) size() int {
	if p == nil {
		return 0
	}
	return 1 + p.left.size() + p.right.size()
}

func (t Tree[T]) Size() int {
	return t.root.size()
}

func (t *Tree[T]) Insert(key T) {
	t.root = insert(t.root, key)
}

func (t *Tree[T]) Erase(key T) {
	t.root = erase(t.root, key)
}

func NewTree[T constraints.Ordered](els ...T) Tree[T] {
	t := Tree[T]{}
	for _, el := range els {
		t.Insert(el)
	}
	return t
}
