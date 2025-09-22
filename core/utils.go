package main

import (
	"fmt"
)

func (n *Number) String() string {
	return fmt.Sprintf("%.2f", n.Value)
}

func (i *Identifier) String() string {
	return i.Name
}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left, b.Op, b.Right)
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("%s%s", u.Op, u.Right)
}

func (p *ParenExpr) String() string {
	return fmt.Sprintf("(%s)", p.Inner)
}

func (a *AssignmentStmt) String() string {
	return fmt.Sprintf("%s = %s", a.Name, a.Value)
}
