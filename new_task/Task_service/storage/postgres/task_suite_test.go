package postgres

import (
	"github/Services/newpro/Task_service/config"
	pb "github/Services/newpro/Task_service/genproto/task_service"
	"github/Services/newpro/Task_service/pkg/db"
	"github/Services/newpro/Task_service/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaskRepoTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepoTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDbForSuite(config.Load())
	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanUpFunc = cleanup
}

func (suite *TaskRepoTestSuite) TestTaskCRUD() {
	task := pb.Task{
		Id:         "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb62",
		Name:       "garry1",
		AssigneeId: "fe66a4bb-4c19-40a1-a5a2-ae8b3eb0cb82",
		Title:      "NewTak3",
		Summary:    "Task3",
		Deadline:   "2022-02-01 05:00",
		Status:     "test",
	}
	_, err := suite.Repository.Delete(&pb.IdReq{Id: task.Id})
	suite.Nil(err)

	task1, err := suite.Repository.Create(&task)
	suite.Nil(err)
	suite.NotNil(task1)

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask, "Task must not be nil")
	suite.Equal(task.AssigneeId, getTask.AssigneeId, "Asignees must match")

	task.Title = "New title"
	updatedTask, err := suite.Repository.Update(&task)
	suite.Nil(err)
	suite.Equal(updatedTask.Title, task.Title, "Titles must match")

	listTasks, err := suite.Repository.List(&pb.ListReq{
		Page:  1,
		Limit: 10,
	})
	suite.Nil(err)
	suite.NotNil(listTasks)

	listOverdueTasks, err := suite.Repository.ListOverdue(&pb.ListOverReq{
		Page:  1,
		Limit: 10,
		Time:  "2022-02-01 04:00",
	})
	suite.Nil(err)
	suite.NotNil(listOverdueTasks)
}

func (suite *TaskRepoTestSuite) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
}
