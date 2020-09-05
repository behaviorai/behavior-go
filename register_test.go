package behavior

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testNode struct {
	BaseNode
}

func TestRegister(t *testing.T) {
	Register("testNode", &testNode{})

	n, err := createNodeByName("testNode")
	assert.Nil(t, err, "err not nil")
	assert.NotNil(t, n, "test node is nil")
	ty := reflect.TypeOf(n).Elem().String()
	assert.Equal(t, "behavior.testNode", ty, "node type should be 'behavior.testNode'")
}
