// package postgres

// import (
// 	"testing"

// 	"github.com/Golang_bot/user-service/config"
// 	pb "github.com/Golang_bot/user-service/genproto"
// 	"github.com/Golang_bot/user-service/pkg/db"
// 	"github.com/Golang_bot/user-service/storage/repo"
// 	"github.com/stretchr/testify/suite"
// )

// type UserRepositoryTestSuite struct {
// 	suite.Suite
// 	CleanupFunc func()
// 	Repository  repo.UserStorageI
// }

// func (suite *UserRepositoryTestSuite) SetupSuite() {
// 	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

// 	suite.Repository = NewUserRepo(pgPool)
// 	suite.CleanupFunc = cleanup
// }

// // All methods that begin with "Test" are run as tests within a
// // suite.
// func (suite *UserRepositoryTestSuite) TestUserCRUD() {
// 	user1 := pb.User{}
// 	user1.FirstName = "salom"
// 	user1.LastName = "hair"
// 	user, err := suite.Repository.CreateUser(&user1)
// 	suite.Nil(err)

// 	getUser, err := suite.Repository.GetUserById(user.Id)
// 	suite.Nil(err)
// 	suite.NotNil(getUser.FirstName, "user must not be nil")
// }

// func (suite *UserRepositoryTestSuite) TearDownSuite() {
// 	suite.CleanupFunc()
// }

// // In order for 'go test' to run this suite, we need to create
// // a normal test function and pass our suite to suite.Run
// func TestUserRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserRepositoryTestSuite))
// }
package postgres 