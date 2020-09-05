package action

import (
	"testing"

	"github.com/billyplus/behavior/config"

	"github.com/stretchr/testify/suite"
)

type actionTestSuite struct {
	suite.Suite
}

func (suite *actionTestSuite) SetupSuite() {
	cfg := &config.BH3Node{
		Properties: map[string]interface{}{
			"duration": int64(500),
		},
	}
	wn := &Wait{}
	err := wn.Initialize(cfg)
	suite.NotNil(err, "initialize wait")

}

func TestWait(t *testing.T) {
	suite.Run(t, &actionTestSuite{})
}
