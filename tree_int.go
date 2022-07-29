package avl

type nodeInt struct {
	left, right *nodeInt
	key         int
	height      int
}

func (p *nodeInt) getHeight() int {
	if p == nil {
		return 0
	}
	return p.height
}

func (p *nodeInt) balanceFactor() int {
	return p.right.getHeight() - p.left.getHeight()
}

func (p *nodeInt) fix() {
	p.height = max(p.left.getHeight(), p.right.getHeight())
}

func rotateLeftInt(p *nodeInt) *nodeInt {
	ret := p.right
	p.right = ret.left
	ret.left = p
	p.fix()
	ret.fix()
	return ret
}

func rotateRightInt(p *nodeInt) *nodeInt {
	ret := p.left
	p.left = ret.right
	ret.right = p
	p.fix()
	ret.fix()
	return ret
}

func balanceInt(p *nodeInt) *nodeInt {
	p.fix()
	switch p.balanceFactor() {
	case -2:
		if p.left.balanceFactor() > 0 {
			p.left = rotateLeftInt(p.left)
		}
		return rotateRightInt(p)
	case 2:
		if p.right.balanceFactor() < 0 {
			p.right = rotateRightInt(p.right)
		}
		return rotateLeftInt(p)
	default:
		return p
	}
}

func insertInt(p *nodeInt, key int) *nodeInt {
	if p == nil {
		return &nodeInt{
			key: key,
		}
	}
	if key < p.key {
		p.left = insertInt(p.left, key)
	} else {
		p.right = insertInt(p.right, key)
	}
	return balanceInt(p)
}

func findMinInt(p *nodeInt) *nodeInt {
	if p.left == nil {
		return p
	}
	return findMinInt(p.left)
}

func eraseMinInt(p *nodeInt) *nodeInt {
	if p.left == nil {
		return p.right
	}
	p.left = eraseMinInt(p.left)
	return balanceInt(p)
}

func eraseInt(p *nodeInt, key int) *nodeInt {
	if p == nil {
		return nil
	}
	if key < p.key {
		p.left = eraseInt(p.left, key)
	} else if key > p.key {
		p.right = eraseInt(p.right, key)
	} else {
		l, r := p.left, p.right
		if r == nil {
			return l
		}
		p = findMinInt(r)
		p.right = eraseMinInt(r)
		p.left = l
	}
	return balanceInt(p)
}

type TreeInt struct {
	root *nodeInt
}

func (t TreeInt) Height() int {
	return t.root.getHeight()
}

func (p *nodeInt) size() int {
	if p == nil {
		return 0
	}
	return 1 + p.left.size() + p.right.size()
}

func (t *TreeInt) Size() int {
	return t.root.size()
}

func (t *TreeInt) Insert(key int) {
	t.root = insertInt(t.root, key)
}

func (t *TreeInt) Erase(key int) {
	t.root = eraseInt(t.root, key)
}

func NewTreeInt(els ...int) TreeInt {
	t := TreeInt{}
	for _, el := range els {
		t.Insert(el)
	}
	return t
}
