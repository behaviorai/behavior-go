package behavior

func newBlackboard(nodeCount int) *Blackboard {
	bb := &Blackboard{
		nodeMemo: make([]Memory, nodeCount),
	}
	// for i := range bb.nodeMemo {
	// 	bb.nodeMemo[i] = make(map[string]interface{})
	// }
	return bb
}

type Blackboard struct {
	Instance    interface{}
	TreeMemo Memory
	nodeMemo []Memory
}

func (bb *Blackboard) GetNodeMemo(index int) Memory {
	return bb.nodeMemo[index]
}

type Memory interface {
	GetStatus() NodeStatus
	SaveStatus(st NodeStatus)
}

type BaseMemo struct {
	status NodeStatus
}

// GetStatus get status of related node
func (memo *BaseMemo) GetStatus() NodeStatus {

	return memo.status
}

// SaveStatus save status of related node
func (memo *BaseMemo) SaveStatus(st NodeStatus) {
	memo.status = st
}

// type Memory map[string]interface{}

// func (memo *Memory) Set(key string, value interface{}) {
// 	(*memo)[key] = value
// }

// func (memo *Memory) Get(key string) interface{} {
// 	return (*memo)[key]
// }

// func (memo *Memory) GetInt(key string) int {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(int)
// }

// func (memo *Memory) GetInt16(key string) int16 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(int16)
// }

// func (memo *Memory) GetInt32(key string) int32 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(int32)
// }

// func (memo *Memory) GetInt64(key string) int64 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(int64)
// }

// func (memo *Memory) GetUint16(key string) uint16 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(uint16)
// }

// func (memo *Memory) GetUint32(key string) uint32 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(uint32)
// }

// func (memo *Memory) GetUint64(key string) uint64 {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(uint64)
// }

// func (memo *Memory) GetBool(key string) bool {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return false
// 	}
// 	return v.(bool)
// }

// func (memo *Memory) GetString(key string) string {
// 	v := (*memo)[key]
// 	if v == nil {
// 		return ""
// 	}
// 	return v.(string)
// }
