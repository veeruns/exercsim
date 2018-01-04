package tree

import (
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

type Mismatch struct{}
type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }
func (r *Record) String() string {
	return fmt.Sprintf("(ID: %d, Parent: %d)", r.ID, r.Parent)
}
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		if r == 0 && (rec.ID != 0 || rec.Parent != 0) {
			return nil, fmt.Errorf("Not a valid root record %s", rec.String())
		} else if r == 0 {
			continue
		} else if r != rec.ID || rec.ID <= rec.Parent {
			return nil, fmt.Errorf("Not a valid record: %s", rec.String())
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}

/*
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
						fmt.Printf("ChildId greater than ParentID\n")
						return nil, errors.New("ChildID greater than Parent ID")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							fmt.Printf("Root ID is not equal to zero %d\n", r.ID)
							return nil, fmt.Errorf("RootID is not equal to zero")
						}
					} else {
						n++
						switch len(c.Children) {
						case 0:
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								c.Children = []*Node{nn, c.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
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
		return nil, Mismatch{}
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}*/
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
