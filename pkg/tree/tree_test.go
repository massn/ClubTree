package tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestUsers struct {
	Max    *User
	Paul   *User
	Niels  *User
	Werner *User
	Erwin  *User
}

func makeTestClubTree(tu TestUsers) (*User, error) {
	r := NewEmptyRoot()
	if err := r.AddUser(tu.Max); err != nil {
		return &User{}, err
	}
	if err := r.AddUser(tu.Paul); err != nil {
		return &User{}, err
	}
	if err := tu.Max.AddUser(tu.Niels); err != nil {
		return &User{}, err
	}
	if err := tu.Niels.AddUser(tu.Werner); err != nil {
		return &User{}, err
	}
	if err := tu.Niels.AddUser(tu.Erwin); err != nil {
		return &User{}, err
	}
	return r, nil
}

func makeTestUsers() TestUsers {
	return TestUsers{
		Max:    makeTestUser("max", "Max", "Planck", ""),
		Paul:   makeTestUser("paul", "Paul", "Dirac", ""),
		Niels:  makeTestUser("niels", "Niels", "Bohr", "max"),
		Werner: makeTestUser("werner", "Werner", "Heisenberg", "niels"),
		Erwin:  makeTestUser("erwin", "Erwin", "Schrodinger", "niels"),
	}
}

func makeTestUser(clubhouseId, firstName, lastName, nominatorId string) *User {
	u := User{}
	u.InitName(&u, clubhouseId)
	u.ClubhouseId = clubhouseId
	u.FirstName = firstName
	u.LastName = lastName
	u.NominatorId = nominatorId
	return &u
}

type CreateTreeTestSuite struct {
	suite.Suite
	TestUsers
	root *User
}

func (suite *CreateTreeTestSuite) SetupTest() {
	suite.TestUsers = makeTestUsers()
	r, err := makeTestClubTree(suite.TestUsers)
	if err != nil {
		panic(err)
	}
	suite.root = r
}

func TestCreateTreeSuite(t *testing.T) {
	suite.Run(t, new(CreateTreeTestSuite))
}

func (suite *CreateTreeTestSuite) TestRootKids() {
	ks := suite.root.Kids
	suite.Assertions.Contains(ks, suite.TestUsers.Max)
}
