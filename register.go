package behavior

import (
	"reflect"
	// "y3pp/log"

	"github.com/pkg/errors"
)

type nodeTypeMap map[string]reflect.Type

var (
	typeMap = make(map[string]reflect.Type)
)

// Register 注册一个新的node类型
//
// 示例： Register("ANode", &ANode{})
func Register(name string, value Node) {
	typeMap[name] = reflect.TypeOf(value).Elem()
}

func createNodeByName(name string) (Node, error) {
	// log.Logger.Debug().Str("name", name).Msg("createNodeByName")
	if nt, ok := typeMap[name]; ok {
		n := reflect.New(nt).Interface()
		// log.Logger.Debug().Str("type", nt.Name()).Interface("value", n).Msg("createNodeByName")
		return n.(Node), nil
	}
	return nil, errors.Errorf("node type not found for '%s'", name)
}
