package node

import (
	"strconv"
)

type Node struct {
	Name     string
	Type     string
	Children interface{}
}

const (
	StringType = "string"
	FloatType  = "float"
	BoolType   = "bool"
	ListType   = "list"
	StructType = "struct"
)

// Structure2Node xxx
func Structure2Node(key string, value interface{}) Node {
	var node Node
	switch value.(type) {
	case bool:
		node.Name = key
		node.Type = BoolType
		node.Children = value
	case string:
		node.Name = key
		node.Type = StringType
		node.Children = value
		break
	case float64:
		node.Name = key
		node.Type = FloatType
		node.Children = value
		break
	case []interface{}:
		data := value.([]interface{})
		node.Name = key
		node.Type = ListType
		var children []Node
		for idx, child := range data {
			children = append(children, Structure2Node(key+"_child_"+strconv.Itoa(idx), child))
		}
		node.Children = children
		break
	case map[string]interface{}:
		data := value.(map[string]interface{})
		node.Name = key
		node.Type = StructType
		children := map[string]Node{}
		for childName, child := range data {
			children[childName] = Structure2Node(childName, child)
		}
		node.Children = children
		break
	}
	return node
}
