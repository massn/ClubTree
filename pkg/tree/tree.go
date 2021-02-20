package tree

import (
	"fmt"
	"os"

	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

var KiT_TypeName = kit.Types.AddType(&User{}, nil)

type ClubTree struct {
	Name string
	Tree *User // Root User
}

type User struct {
	ki.Node
	ClubhouseId string
	Alias       string
	NominatorId string
	FirstName   string
	LastName    string
	TwitterId   string
	InstagramId string
}

func NewClubTree(name string) *ClubTree {
	e := User{}
	e.InitName(&e, "EmptyRoot")
	return &ClubTree{
		Name: name,
		Tree: &e,
	}
}

func (ct *ClubTree) AddUser(u *User) error {
	root := ct.Tree
	c := root.ChildByName(u.Name(), -1)
	if c != nil {
		return fmt.Errorf("The same name user already exists.")
	}

	connectedToChild := false
	for _, child := range getAllChildren(root) {
		if u.ClubhouseId == child.NominatorId {
			ki.MoveToParent(child, u)
		}
		if u.NominatorId == child.ClubhouseId {
			err := child.AddChild(u)
			if err != nil {
				return err
			}
			connectedToChild = true
		}
	}

	if !connectedToChild {
		return root.AddChild(u)
	}
	return nil
}

func getAllChildren(root *User) []*User {
	children := []*User{root}
	for _, c := range *(root.Children()) {
		children = append(children, getAllChildren(c.(*User))...)
	}
	return children
}

func NewDummyUser(id, nominatorId string) *User {
	u := User{}
	u.InitName(&u, id)
	u.ClubhouseId = id
	u.NominatorId = nominatorId
	return &u
}

func NewUser() *User {
	clubhouseId := readInputString("Enter the clubhouse ID", true)
	u := User{}
	u.InitName(&u, clubhouseId)
	u.ClubhouseId = clubhouseId
	u.Alias = readInputString("Enter the alias", false)
	u.FirstName = readInputString("Enter the first name", true)
	u.LastName = readInputString("Enter the last name", true)
	u.TwitterId = readInputString("Enter the Twitter ID", false)
	u.InstagramId = readInputString("Enter the Instagram ID", false)
	u.NominatorId = readInputString("Enter the nominator ID", false)
	return &u
}

func readInputString(message string, required bool) string {
	input := ""
	for {
		fmt.Println(message)
		fmt.Scanf("%s", &input)
		if required && input == "" {
			fmt.Println("Empty input is not allowed.")
			continue
		}
		break
	}
	return input
}

func ReadJson(filename string) (*ClubTree, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &ClubTree{}, err
	}
	defer f.Close()
	ct := NewClubTree(filename)
	if err := ct.Tree.ReadJSON(f); err != nil {
		return &ClubTree{}, err
	}
	return ct, nil
}
