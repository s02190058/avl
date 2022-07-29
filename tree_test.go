package avl

import (
	"math"
	"math/rand"
	"testing"
)

func TestTree(t *testing.T) {
	// n belongs to [a, b)
	n := a + rand.Intn(b-a)
	els := randIntSlice(n)
	tree := NewTree(els...)

	size := tree.Size()
	if size != n {
		t.Errorf("expected size: %d, got: %d", n, size)
	}
	t.Logf("size: %d", size)
	if !tree.isSorted() {
		t.Error("tree isn't sorted")
	}
	if !tree.isBalanced() {
		t.Error("tree isn't balanced")
	}
	// max_h := A * log_2(n + 2) - B
	maxHeight := int(math.Ceil(A*math.Log2(float64(n+2)) - B))
	height := tree.Height()
	if height > maxHeight {
		t.Errorf("expected max height: %d, got: %d", maxHeight, height)
	}
	t.Logf("height: %d", height)

	// erasing half of the elements
	half := n / 2
	for i := 0; i < half; i++ {
		tree.Erase(els[i])
	}
	els = els[half:]
	n -= half

	size = tree.Size()
	if tree.Size() != n {
		t.Errorf("expected size: %d, got: %d", n, size)
	}
	t.Logf("size: %d", size)
	if !tree.isSorted() {
		t.Error("tree isn't sorted")
	}
	if !tree.isBalanced() {
		t.Error("tree isn't balanced")
	}
	// max_h := A * log_2(n + 2) - B
	maxHeight = int(math.Ceil(A*math.Log2(float64(n+2)) - B))
	height = tree.Height()
	if height > maxHeight {
		t.Errorf("expected max height: %d, got: %d", maxHeight, height)
	}
	t.Logf("height: %d", height)

	// erasing the rest
	for i := 0; i < n; i++ {
		tree.Erase(els[i])
	}
	size = tree.Size()
	if size != 0 {
		t.Errorf("expected size: 0, got: %d", size)
	}
	if !tree.isSorted() {
		t.Error("tree isn't sorted")
	}
	if !tree.isBalanced() {
		t.Error("tree isn't balanced")
	}
	height = tree.Height()
	if height != 0 {
		t.Errorf("expected max height: 0, got: %d", height)
	}
}

func (p *node[T]) isSorted() bool {
	if p == nil {
		return true
	}
	if p.left != nil && p.left.key >= p.key {
		return false
	}
	if p.right != nil && p.right.key < p.key {
		return false
	}
	return p.left.isSorted() && p.right.isSorted()
}

func (t *Tree[T]) isSorted() bool {
	return t.root.isSorted()
}

func (p *node[T]) isBalanced() bool {
	if p == nil {
		return true
	}
	bf := p.balanceFactor()
	if bf < -1 || bf > 1 {
		return false
	}
	return p.left.isBalanced() && p.right.isBalanced()
}

func (t *Tree[T]) isBalanced() bool {
	return t.root.isBalanced()
}

func BenchmarkGenTreeInt(b *testing.B) {
	benchCases := []struct {
		name string
	}{
		{
			name: "32",
		},
		{
			name: "1024",
		},
		{
			name: "32768",
		},
	}

	for idx, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tree := NewTree(dataSetsInt[idx]...)
				for _, el := range dataSetsInt[idx] {
					tree.Erase(el)
				}
			}
		})
	}
}

func BenchmarkGenTreeString(b *testing.B) {
	benchCases := []struct {
		name string
	}{
		{
			name: "32",
		},
		{
			name: "1024",
		},
		{
			name: "32768",
		},
	}

	for idx, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tree := NewTree(dataSetsString[idx]...)
				for _, el := range dataSetsString[idx] {
					tree.Erase(el)
				}
			}
		})
	}
}
