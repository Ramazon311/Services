package postgres

import (
	"github/Services/newpro/User_service/config"
	pb "github/Services/newpro/User_service/genproto/user_service"
	"github/Services/newpro/User_service/pkg/db"
	"github/Services/newpro/User_service/storage/repo"

	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.UserStorageI
}

func (suite *UserRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewUserRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

func (suite *UserRepoTestSuite) TestUserCRUD() {
	User := pb.User{
		Id:           "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		FirstName:    "garry1",
		LastName:     "fe66a4bb",
		Username:     "Newuser",
		ProfilePhoto: "User3",
		Bio:          "hbdkvguer8ofgo",
		Email:        "test",
		Gender:       "test",
		Address:      "test",
		PhoneNumber:  "test",
	}
	_, err := suite.Repository.Delete(&pb.IdReq{Id: User.Id})
	suite.Nil(err)

	User1, err := suite.Repository.Create(&User)
	suite.Nil(err)
	suite.NotNil(User1)

	getUser, err := suite.Repository.Get(User.Id)
	suite.Nil(err)
	suite.NotNil(getUser, "User must not be nil")
	suite.Equal(User.Id, getUser.Id, "Asignees must match")

	User.Username = "New user"
	updatedUser, err := suite.Repository.Update(&User)
	suite.Nil(err)
	suite.Equal(updatedUser.Username, User.Username, "Titles must match")

	listUsers, err := suite.Repository.List(&pb.ListReq{
		Page:  1,
		Limit: 10,
	})
	suite.Nil(err)
	suite.NotNil(listUsers)
}

func (suite *UserRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
