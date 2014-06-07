package symbol

import (
	"sort"
)

type order []Symbol

func (o order) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o order) Len() int {
	return len(o)
}

func (o order) Less(i, j int) bool {
	a := o[i]
	b := o[j]

	ta := a.Type()
	tb := b.Type()
	if ta < tb {
		return true
	}
	if ta > tb {
		return false
	}

	return a.Ident() < b.Ident()
}

func Sort(syms []Symbol) {
	sort.Sort(order(syms))
}
