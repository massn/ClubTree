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
	Albert *User
}

func makeTestClubTree(tu TestUsers) (*ClubTree, error) {
	ct := NewClubTree("test_club_tree")
	if err := ct.AddUser(tu.Max); err != nil {
		return &ClubTree{}, err
	}
	if err := ct.AddUser(tu.Paul); err != nil {
		return &ClubTree{}, err
	}
	if err := ct.AddUser(tu.Niels); err != nil {
		return &ClubTree{}, err
	}
	if err := ct.AddUser(tu.Werner); err != nil {
		return &ClubTree{}, err
	}
	if err := ct.AddUser(tu.Erwin); err != nil {
		return &ClubTree{}, err
	}
	return ct, nil
}

func makeTestUsers() TestUsers {
	return TestUsers{
		Max:    makeTestUser("max", "Max", "Planck", "gustav"),
		Paul:   makeTestUser("paul", "Paul", "Dirac", "ralph"),
		Niels:  makeTestUser("niels", "Niels", "Bohr", "max"),
		Werner: makeTestUser("werner", "Werner", "Heisenberg", "niels"),
		Erwin:  makeTestUser("erwin", "Erwin", "Schrodinger", "niels"),
		Albert: makeTestUser("albert", "Albert", "Einstein", "albert"),
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
	clubTree *ClubTree
}

func (suite *CreateTreeTestSuite) SetupTest() {
	suite.TestUsers = makeTestUsers()
	r, err := makeTestClubTree(suite.TestUsers)
	if err != nil {
		panic(err)
	}
	suite.clubTree = r
}

func TestCreateTreeSuite(t *testing.T) {
	suite.Run(t, new(CreateTreeTestSuite))
}

func (suite *CreateTreeTestSuite) TestRootChildren() {
	ks := suite.clubTree.Tree.Children()
	suite.Assertions.Equal(len(*ks), 2)
	suite.Assertions.Contains(*ks, suite.TestUsers.Max)
	suite.Assertions.Contains(*ks, suite.TestUsers.Paul)
}

func (suite *CreateTreeTestSuite) TestChildrenOfMax() {
	ks := suite.TestUsers.Max.Children()
	suite.Assertions.Equal(len(*ks), 1)
	suite.Assertions.Contains(*ks, suite.TestUsers.Niels)
}

func (suite *CreateTreeTestSuite) TestChildrenOfNiels() {
	ks := suite.TestUsers.Niels.Children()
	suite.Assertions.Equal(len(*ks), 2)
	suite.Assertions.Contains(*ks, suite.TestUsers.Werner)
	suite.Assertions.Contains(*ks, suite.TestUsers.Erwin)
}

func (suite *CreateTreeTestSuite) TestChildrenOfWerner() {
	ks := suite.TestUsers.Werner.Children()
	suite.Assertions.Equal(len(*ks), 0)
}

func (suite *CreateTreeTestSuite) TestChildrenOfErwin() {
	ks := suite.TestUsers.Erwin.Children()
	suite.Assertions.Equal(len(*ks), 0)
}

type AddTreeTestSuite struct {
	suite.Suite
	TestUsers
	clubTree *ClubTree
}

func (suite *AddTreeTestSuite) SetupTest() {
	suite.TestUsers = makeTestUsers()
	r, err := makeTestClubTree(suite.TestUsers)
	if err != nil {
		panic(err)
	}
	suite.clubTree = r
}

func TestAddTreeSuite(t *testing.T) {
	suite.Run(t, new(AddTreeTestSuite))
}

func (suite *AddTreeTestSuite) TestAddToRoot() {
	err := suite.clubTree.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.clubTree.Tree.Children()
	suite.Assertions.Equal(len(*ks), 3)
	suite.Assertions.Contains(*ks, suite.TestUsers.Albert)
}

func (suite *AddTreeTestSuite) TestAddToPaul() {
	suite.TestUsers.Albert.NominatorId = "paul"
	err := suite.clubTree.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.clubTree.Tree.Children()
	suite.Assertions.Equal(len(*ks), 2)
	paulChildren := suite.TestUsers.Paul.Children()
	suite.Assertions.Equal(len(*paulChildren), 1)
	suite.Assertions.Contains(*paulChildren, suite.TestUsers.Albert)
}

func (suite *AddTreeTestSuite) TestAddToErwin() {
	suite.TestUsers.Albert.NominatorId = "erwin"
	err := suite.clubTree.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.clubTree.Tree.Children()
	suite.Assertions.Equal(len(*ks), 2)
	maxChildren := suite.TestUsers.Max.Children()
	suite.Assertions.Equal(len(*maxChildren), 1)
	nielsChildren := suite.TestUsers.Niels.Children()
	suite.Assertions.Equal(len(*nielsChildren), 2)
	wernerChildren := suite.TestUsers.Werner.Children()
	suite.Assertions.Equal(len(*wernerChildren), 0)
	erwinChildren := suite.TestUsers.Erwin.Children()
	suite.Assertions.Equal(len(*erwinChildren), 1)
	suite.Assertions.Contains(*erwinChildren, suite.TestUsers.Albert)
}
