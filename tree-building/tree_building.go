package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Node struct {
	ID       int
	Children []*Node
}

func Build2(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(byID(records))

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("Root not found")
	}

	if len(records) == 1 {
		return &Node{ID: 0}, nil
	}

	children, er := getChildren(0, records[1:])
	if er != nil {
		return nil, er
	}

	for i := 1; i < len(records); i++ {
		rec := records[i]
		if rec.ID != i {
			return nil, errors.New("Non-continuous or duplicated IDs")
		}
		if rec.ID > 0 && rec.ID <= rec.Parent {
			return nil, errors.New("ID higher or equals than parent")
		}
	}

	return &Node{ID: 0, Children: children}, nil
}

func getChildren(parent int, records []Record) ([]*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	siblings := []Record{}
	children := []Record{}
	for _, r := range records {
		if r.Parent == parent {
			siblings = append(siblings, r)
		} else {
			children = append(children, r)
		}
	}

	result := []*Node{}
	for _, r := range siblings {
		nodeChildren, err := getChildren(r.ID, children)
		if err != nil {
			return nil, err
		}
		result = append(result, &Node{ID: r.ID, Children: nodeChildren})
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func removeItem(slice []Record, idx int) []Record {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := &Node{}
	todo := []*Node{root}
	n := 1
	for {
		if len(todo) == 0 {
			break
		}
		newTodo := []*Node(nil)
		for _, c := range todo {
			for _, r := range records {
				if r.Parent == c.ID {
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("b")
						}
					} else {
						n++
						nn := &Node{ID: r.ID}
						newTodo = append(newTodo, nn)
						switch len(c.Children) {
						case 0:
							c.Children = []*Node{nn}
						case 1:
							if c.Children[0].ID == r.ID {
								return nil, errors.New("Duplicated ID")
							} else if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
							} else {
								c.Children = []*Node{nn, c.Children[0]}
							}
						default:
						breakpoint:
							for range []bool{false} {
								for i, cc := range c.Children {
									if cc.ID > r.ID {
										a := make([]*Node, len(c.Children)+1)
										copy(a, c.Children[:i])
										copy(a[i+1:], c.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										c.Children = a
										break breakpoint
									}
								}
								c.Children = append(c.Children, nn)
							}
						}
					}
				}
			}
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, errors.New("c")
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
