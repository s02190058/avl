package avl

type nodeString struct {
	left, right *nodeString
	key         string
	height      int
}

func (p *nodeString) getHeight() int {
	if p == nil {
		return 0
	}
	return p.height
}

func (p *nodeString) balanceFactor() int {
	return p.right.getHeight() - p.left.getHeight()
}

func (p *nodeString) fix() {
	p.height = max(p.left.getHeight(), p.right.getHeight())
}

func rotateLeftString(p *nodeString) *nodeString {
	ret := p.right
	p.right = ret.left
	ret.left = p
	p.fix()
	ret.fix()
	return ret
}

func rotateRightString(p *nodeString) *nodeString {
	ret := p.left
	p.left = ret.right
	ret.right = p
	p.fix()
	ret.fix()
	return ret
}

func balanceString(p *nodeString) *nodeString {
	p.fix()
	switch p.balanceFactor() {
	case -2:
		if p.left.balanceFactor() > 0 {
			p.left = rotateLeftString(p.left)
		}
		return rotateRightString(p)
	case 2:
		if p.right.balanceFactor() < 0 {
			p.right = rotateRightString(p.right)
		}
		return rotateLeftString(p)
	default:
		return p
	}
}

func insertString(p *nodeString, key string) *nodeString {
	if p == nil {
		return &nodeString{
			key: key,
		}
	}
	if key < p.key {
		p.left = insertString(p.left, key)
	} else {
		p.right = insertString(p.right, key)
	}
	return balanceString(p)
}

func findMinString(p *nodeString) *nodeString {
	if p.left == nil {
		return p
	}
	return findMinString(p.left)
}

func eraseMinString(p *nodeString) *nodeString {
	if p.left == nil {
		return p.right
	}
	p.left = eraseMinString(p.left)
	return balanceString(p)
}

func eraseString(p *nodeString, key string) *nodeString {
	if p == nil {
		return nil
	}
	if key < p.key {
		p.left = eraseString(p.left, key)
	} else if key > p.key {
		p.right = eraseString(p.right, key)
	} else {
		l, r := p.left, p.right
		if r == nil {
			return l
		}
		p = findMinString(r)
		p.right = eraseMinString(r)
		p.left = l
	}
	return balanceString(p)
}

type TreeString struct {
	root *nodeString
}

func (t TreeString) Height() int {
	return t.root.getHeight()
}

func (p *nodeString) size() int {
	if p == nil {
		return 0
	}
	return 1 + p.left.size() + p.right.size()
}

func (t *TreeString) Size() int {
	return t.root.size()
}

func (t *TreeString) Insert(key string) {
	t.root = insertString(t.root, key)
}

func (t *TreeString) Erase(key string) {
	t.root = eraseString(t.root, key)
}

func NewTreeString(els ...string) TreeString {
	t := TreeString{}
	for _, el := range els {
		t.Insert(el)
	}
	return t
}
