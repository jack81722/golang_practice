package async

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskSuite struct {
	suite.Suite
}

func (t *TaskSuite) TestRunTask() {
	num := 0
	task := NewTask(func() {
		for i := 1; i <= 10; i++ {
			num += i
		}
	})
	task.Start()
	task.Wait()
	assert.Equal(t.T(), 55, num)
}

func (t *TaskSuite) TestContinueWith() {
	num := 0
	sumFunc := func() {
		for i := 1; i <= 10; i++ {
			num += i
		}
	}
	task := NewTask(sumFunc).
		ContinueWith(sumFunc).
		ContinueWith(sumFunc)
	task.Start()
	task.Wait()
	if !assert.Equal(t.T(), 165, num) {
		assert.Fail(t.T(), "not equal")
	}
}

func (t *TaskSuite) TestTaskError() {
	errFunc := func() {
		panic(errors.New("task error"))
	}
	task := NewTask(errFunc)
	task.Start()
	task.Wait()
	assert.Error(t.T(), task.err)
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskSuite))
}
