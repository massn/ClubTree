package main

import (
	"fmt"
	//"os"

	"github.com/goki/ki/ki"
)

type EmptyRoot struct {
	ki.Node
}

func (r *EmptyRoot) AddUser(u *User) error {
	c := r.ChildByName(u.Name(), -1)
	if c != nil {
		return fmt.Errorf("The same name user already exists.")
	}

	for _, child := range *(r.Children()) {
		if child.(*User).NominatorId == u.Name() {
			ki.MoveToParent(child, u)
			break
		}
	}

	nominatorOfUser := r.ChildByName(u.NominatorId, -1)
	if nominatorOfUser == nil {
		return r.AddChild(u)
	}
	return nominatorOfUser.AddChild(u)
}

type User struct {
	ki.Node
	ClubhouseId string
	NominatorId string
	FirstName   string
	LastName    string
	TwitterId   string
	InstagramId string
}

func newEmptyRoot() *EmptyRoot {
	e := EmptyRoot{}
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
	filename := "clubtree.json"
	root := newEmptyRoot()

	newUserId := "user"
	newNominatorId := "nominator"

	user := newDummyUser(newUserId, newNominatorId)
	nominator := newDummyUser("nominator", "nominators-nominator")

	_ = root.AddUser(user)
	_ = root.AddUser(nominator)

	if err := root.SaveJSON(filename); err != nil {
		panic(err)
	}
}
