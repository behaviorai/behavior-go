package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// BH3Node behavior3的节点
type BH3Node struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Children    []string `json:"children"`
	Child       string   `json:"child"`
	// Parameters  map[string]interface{} `json:"parameters"`
	Properties map[string]interface{} `json:"properties"`
}

// BH3Tree behavior3树
type BH3Tree struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Root        string                 `json:"root"`
	Properties  map[string]interface{} `json:"properties"`
	Nodes       map[string]BH3Node     `json:"nodes"`
}

// BH3Project behavior3的工程json类型
type BH3Project struct {
	SelectedTree string    `json:"selectedTree"`
	Scope        string    `json:"scope"`
	Trees        []BH3Tree `json:"trees"`
}

// B3File 是behavior3编辑器保存的.b3格式的配置文件
type B3File struct {
	Name string     `json:"name"`
	Data BH3Project `json:"data"`
}

func LoadB3File(path string) (*B3File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseB3File(data)
}

// ParseB3File parse project from []byte
func ParseB3File(data []byte) (*B3File, error) {
	var proj B3File
	err := json.Unmarshal(data, &proj)
	if err != nil {
		return nil, err
	}
	return &proj, nil
}

// // LoadBH3Project 加载behavior3的project配置格式
// func LoadBH3Project(path string) (*BH3Project, error) {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ParseBH3Project(data)
// }

// // ParseBH3Project parse project from []byte
// func ParseBH3Project(data []byte) (*BH3Project, error) {
// 	var proj BH3Project
// 	err := json.Unmarshal(data, &proj)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &proj, nil
// }

func (node *BH3Node) GetFloat64(name string) float64 {
	v := node.Properties[name]
	if v == nil {
		return 0
	}
	return v.(float64)
}

func (node *BH3Node) GetInt(name string) int {
	return int(node.GetFloat64(name))
}

func (node *BH3Node) GetInt32(name string) int32 {
	return int32(node.GetFloat64(name))
}

func (node *BH3Node) GetInt64(name string) int64 {
	return int64(node.GetFloat64(name))
}

func (node *BH3Node) GetUint32(name string) uint32 {
	return uint32(node.GetFloat64(name))
}

func (node *BH3Node) GetUint64(name string) uint64 {
	return uint64(node.GetFloat64(name))
}

func (node *BH3Node) GetString(name string) string {
	return (node.Properties[name]).(string)
}

func (node *BH3Node) GetBool(name string) bool {
	return (node.Properties[name]).(bool)
}

func (node *BH3Node) GetInt32s(name string) []int32 {
	var ret []int32
	v := (node.Properties[name]).(string)
	if v == "" {
		return nil
	}
	lst := strings.Split(v, ",")
	for _, val := range lst {
		n, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Errorf("failed to unmarshal int32s: %v", err))
		}
		ret = append(ret, int32(n))
	}
	return ret
}

func (node *BH3Node) GetInt64s(name string) []int64 {
	var ret []int64
	v := (node.Properties[name]).(string)
	if v == "" {
		return nil
	}
	lst := strings.Split(v, ",")
	for _, val := range lst {
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to unmarshal int32s: %v", err))
		}
		ret = append(ret, n)
	}
	return ret
}
