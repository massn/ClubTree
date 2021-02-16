package tree

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

func (u *User) Print() {
	t := toGoTree(u)
	fmt.Println(t.Print())
}

func toGoTree(u *User) gotree.Tree {
	root := gotree.New(makeTreeId(u))
	for _, kid := range u.Kids {
		if kid.(*User).Kids == nil {
			root.Add(makeTreeId(kid.(*User)))
		} else {
			root.AddTree(toGoTree(kid.(*User)))
		}
	}
	return root
}

func makeTreeId(u *User) string {
	return fmt.Sprintf("%s/%s %s", u.ClubhouseId, u.FirstName, u.LastName)
}
