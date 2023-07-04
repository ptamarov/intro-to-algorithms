package bsts

type valueGetter interface {
	getValue() int
}

type value int

func (v value) getValue() int {
	return int(v)
}

func BinaryTreeFromList(sl []int) Tree {
	vs := parseList(sl)
	l := len(sl)
	ts := make([](*node), len(sl))

loop:
	for j, s := range vs {
		if s == nil {
			continue loop
		} else {
			ts[j] = &node{}
			ts[j].value = s.getValue()
		}
	}
	for j := range vs {
		if 2*j+1 < l {
			if vs[2*j+1] != nil {
				ts[j].left = ts[2*j+1]
				ts[2*j+1].parent = ts[j]
			}
		}
		if 2*j+2 < l {
			if vs[2*j+2] != nil {
				ts[j].right = ts[2*j+2]
				ts[2*j+2].parent = ts[j]
			}
		}
	}
	return Tree{ts[0]}
}

// parseList takes a list of non-negative integers and -1s and generates
// a list of ValueGetters assigning nil to -1.
func parseList(ls []int) []valueGetter {
	v := make([]valueGetter, 0, len(ls))
	for _, s := range ls {
		if s == -1 {
			v = append(v, nil)
		} else {
			v = append(v, value(s))
		}
	}
	return v
}
