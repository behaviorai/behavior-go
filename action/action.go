package action

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

const (
	tagDuration = "duration"
)

type Action struct {
	behavior.BaseNode
}

func (node *Action) Initialize(cfg *config.BH3Node) error {
	return node.BaseNode.Initialize(cfg)
}

func (node *Action) String() string {
	return "Action"
}
