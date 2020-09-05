package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type bh3TestSuite struct {
	suite.Suite
}

var testnode = []byte(`
{
    "version": "0.3.0",
    "scope": "node",
    "name": "UntilFailure",
    "category": "decorator",
    "title": "UntilFailure",
    "description": "重复执行子节点，直到子节点返回failed，此节点会返回success。或者达到maxLoop次后，此节点返回failed",
    "properties": {
	  "IntV": 100,
	  "Int64": 400,
	  "boolT": true,
	  "boolF": false
    }
}
`)

func (suite *bh3TestSuite) SetupSuite() {
	var ufnode BH3Node
	err := json.Unmarshal(testnode, &ufnode)
	suite.Nil(err, "unmarshal testnode")
	suite.Equal("UntilFailure", ufnode.Name, "check name")
	suite.Equal("UntilFailure", ufnode.Title, "check title")
	suite.Equal("decorator", ufnode.Category, "check category")

	IntV := ufnode.GetInt("IntV")
	suite.Equal(100, IntV, "check IntV")

	Int64 := ufnode.GetInt("Int64")
	suite.Equal(400, Int64, "check Int64")

	boolF := ufnode.GetBool("boolF")
	suite.Equal(false, boolF, "check boolF")

	boolT := ufnode.GetBool("boolT")
	suite.Equal(true, boolT, "check boolT")
}

func TestBehavior3(t *testing.T) {
	suite.Run(t, &bh3TestSuite{})
}
