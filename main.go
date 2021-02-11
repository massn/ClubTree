package main

import (
	"fmt"
	"os"

	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

var KiT_TypeName = kit.Types.AddType(&User{}, nil)

type User struct {
	ki.Node
	ClubhouseId string
	NominatorId string
	FirstName   string
	LastName    string
	TwitterId   string
	InstagramId string
}

func (root *User) AddUser(u *User) error {
	c := root.ChildByName(u.Name(), -1)
	if c != nil {
		return fmt.Errorf("The same name user already exists.")
	}

	for _, child := range *(root.Children()) {
		if child.(*User).NominatorId == u.Name() {
			ki.MoveToParent(child, u)
			break
		}
	}

	nominatorOfUser := root.ChildByName(u.NominatorId, -1)
	if nominatorOfUser == nil {
		return root.AddChild(u)
	}
	return nominatorOfUser.AddChild(u)
}

func newEmptyRoot() *User {
	e := User{}
	e.InitName(&e, "EmptyRoot")
	return &e
}

func newDummyUser(id, nominatorId string) *User {
	u := User{}
	u.InitName(&u, id)
	u.ClubhouseId = id
	u.NominatorId = nominatorId
	return &u
}

func newUser() *User {
	clubhouseId := readInputString("Enter the clubhouse ID", true)
	u := User{}
	u.InitName(&u, clubhouseId)
	u.ClubhouseId = clubhouseId
	u.NominatorId = readInputString("Enter the nominator ID", true)
	u.FirstName = readInputString("Enter the first name", true)
	u.LastName = readInputString("Enter the last name", true)
	u.TwitterId = readInputString("Enter the Twitter ID", false)
	u.InstagramId = readInputString("Enter the Instagram ID", false)
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

func main() {
	oldFilename := "clubtree.json"
	newFilename := "new-clubtree.json"
	root, err := readJson(oldFilename)
	if err != nil {
		panic(err)
	}

	newUserId := "newuser"
	newNominatorId := "olduser"

	user := newDummyUser(newUserId, newNominatorId)

	_ = root.AddUser(user)

	if err := root.SaveJSON(newFilename); err != nil {
		panic(err)
	}
}

func readJson(filename string) (*User, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &User{}, err
	}
	defer f.Close()
	root := newEmptyRoot()
	if err := root.ReadJSON(f); err != nil {
		return &User{}, err
	}
	return root, nil
}
