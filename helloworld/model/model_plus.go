package model

import "fmt"

type Plus struct {
	A int
	B int
}

type TriplePlus struct {
	Plus
	C int
}

func (p *Plus) Cal() int {
	return p.A + p.B
}

func (p *Plus) Print() string {
	return fmt.Sprintf("A=%v, B=%v\n", p.A, p.B)
}

// func (p *TriplePlus) Cal() int {
// 	return p.A + p.B + p.C
// }
