package rx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RxSuite struct {
	suite.Suite
}

func (t *RxSuite) TestJust() {
	j := Just(1)
	s := &Subject{
		NextHandler:     func(i interface{}) { fmt.Println(i) },
		ErrorHandler:    func(err error) { assert.Fail(t.T(), "error") },
		CompleteHandler: func() { fmt.Println("completed") },
	}
	j.Subscribe(s)
}

func TestRxSuiteInit(t *testing.T) {
	suite.Run(t, new(RxSuite))
}
