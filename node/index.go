package node


type Node struct {
	Name string
	Type string
	Children interface{}
}

// Structure2Node xxx
func Structure2Node(key string, value interface{}) Node {
	var node Node
	switch value.(type) {
	case string:
		node.Name = key
		node.Type = "string"
		node.Children = value
	case float64:
		node.Name = key
		node.Type = "float64"
		node.Children = value
	case []interface{}:
		data := value.([]interface{})
		node.Name = key
		node.Type = "[]-any"
		var children []Node
		for idx, child := range data {
			children = append(children, Structure2Node("child_" + string(idx), child))
		}
		node.Children = children
	case map[string]interface{}:
		data := value.(map[string]interface{})
		node.Name = key
		node.Type = "struct"
		children := map[string]Node{}
		for childName, child := range data {
			children[childName] = Structure2Node(childName, child)
		}
		node.Children = children
	}
	return node
}
