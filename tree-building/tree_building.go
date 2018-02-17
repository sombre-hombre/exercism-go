package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	recordsLength := len(records)
	if recordsLength == 0 {
		return nil, nil
	}

	var nodes = make(map[int]*Node)
	for _, record := range records {
		if _, found := nodes[record.ID]; found {
			return nil, fmt.Errorf("Duplicate ID %d found", record.ID)
		}
		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, errors.New("Root node has parent")
			}
		} else {
			if record.ID == record.Parent {
				return nil, fmt.Errorf("Self parent node %d", record.ID)
			}
			if record.ID <= record.Parent {
				return nil, fmt.Errorf("Wrong link %d → %d", record.ID, record.Parent)
			}
		}

		nodes[record.ID] = &Node{ID: record.ID}
	}

	if _, found := nodes[0]; !found {
		return nil, errors.New("Root node not found")
	}
	if last, found := nodes[recordsLength-1]; !found || last.ID != recordsLength-1 {
		return nil, errors.New("Non-continuous records")
	}

	for _, record := range records {
		node := nodes[record.Parent]

		if record.ID != 0 {
			node.Children = append(node.Children, nodes[record.ID])
		}
	}

	for i := 0; i < recordsLength; i++ {
		if len(nodes[i].Children) > 1 {
			sort.Sort(byID(nodes[i].Children))
		}
	}

	return nodes[0], nil
}

type byID []*Node

func (arr byID) Len() int {
	return len(arr)
}

func (arr byID) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr byID) Less(i, j int) bool {
	return arr[i].ID < arr[j].ID
}
