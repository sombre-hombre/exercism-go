package sublist

type Relation string

const (
	EQUAL     Relation = "equal"
	SUBLIST   Relation = "sublist"
	SUPERLIST Relation = "superlist"
	UNEQUAL   Relation = "unequal"
)

func Sublist(a, b intArray) Relation {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return UNEQUAL
			}
		}
		return EQUAL
	}

	if len(a) < len(b) {
		if b.contains(a) {
			return SUBLIST
		}
		return UNEQUAL
	}

	if a.contains(b) {
		return SUPERLIST
	}
	return UNEQUAL
}

type intArray []int

func (a intArray) contains(b []int) bool {
	if len(a) < len(b) {
		return false
	}
	if len(b) == 0 {
		return true
	}

	var offset, i int
	for i < len(b) && i+offset < len(a) {
		if b[i] != a[i+offset] {
			offset++
			i = 0
			continue
		}
		i++
	}

	return i == len(b)
}
