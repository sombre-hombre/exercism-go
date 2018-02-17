package tree

type Record struct {
	ID, Parent int
}

type byID []Record

func (arr byID) Len() int {
	return len(arr)
}

func (arr byID) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr byID) Less(i, j int) bool {
	return arr[i].ID < arr[j].ID
}
