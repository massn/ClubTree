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
	suite.Assertions.Equal(len(ks), 2)
	suite.Assertions.Contains(ks, suite.TestUsers.Max)
	suite.Assertions.Contains(ks, suite.TestUsers.Paul)
}

func (suite *CreateTreeTestSuite) TestKidsOfMax() {
	ks := suite.TestUsers.Max.Kids
	suite.Assertions.Equal(len(ks), 1)
	suite.Assertions.Contains(ks, suite.TestUsers.Niels)
}

func (suite *CreateTreeTestSuite) TestKidsOfNiels() {
	ks := suite.TestUsers.Niels.Kids
	suite.Assertions.Equal(len(ks), 2)
	suite.Assertions.Contains(ks, suite.TestUsers.Werner)
	suite.Assertions.Contains(ks, suite.TestUsers.Erwin)
}

func (suite *CreateTreeTestSuite) TestKidsOfWerner() {
	ks := suite.TestUsers.Werner.Kids
	suite.Assertions.Equal(len(ks), 0)
}

func (suite *CreateTreeTestSuite) TestKidsOfErwin() {
	ks := suite.TestUsers.Erwin.Kids
	suite.Assertions.Equal(len(ks), 0)
}

type AddTreeTestSuite struct {
	suite.Suite
	TestUsers
	root *User
}

func (suite *AddTreeTestSuite) SetupTest() {
	suite.TestUsers = makeTestUsers()
	r, err := makeTestClubTree(suite.TestUsers)
	if err != nil {
		panic(err)
	}
	suite.root = r
}

func TestAddTreeSuite(t *testing.T) {
	suite.Run(t, new(AddTreeTestSuite))
}

func (suite *AddTreeTestSuite) TestAddToRoot() {
	err := suite.root.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.root.Kids
	suite.Assertions.Equal(len(ks), 3)
	suite.Assertions.Contains(ks, suite.TestUsers.Albert)
}

func (suite *AddTreeTestSuite) TestAddToPaul() {
	err := suite.TestUsers.Paul.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.root.Kids
	suite.Assertions.Equal(len(ks), 2)
	paulKids := suite.TestUsers.Paul.Kids
	suite.Assertions.Equal(len(paulKids), 1)
	suite.Assertions.Contains(paulKids, suite.TestUsers.Albert)
}

func (suite *AddTreeTestSuite) TestAddToErwin() {
	err := suite.TestUsers.Erwin.AddUser(suite.TestUsers.Albert)
	suite.Assertions.Nil(err)
	ks := suite.root.Kids
	suite.Assertions.Equal(len(ks), 2)
	maxKids := suite.TestUsers.Max.Kids
	suite.Assertions.Equal(len(maxKids), 1)
	nielsKids := suite.TestUsers.Niels.Kids
	suite.Assertions.Equal(len(nielsKids), 2)
	wernerKids := suite.TestUsers.Werner.Kids
	suite.Assertions.Equal(len(wernerKids), 0)
	erwinKids := suite.TestUsers.Erwin.Kids
	suite.Assertions.Equal(len(erwinKids), 1)
	suite.Assertions.Contains(erwinKids, suite.TestUsers.Albert)
}
