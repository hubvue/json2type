package generator

import (
	"fmt"
	"github.com/hubvue/json2type/node"
	"strings"
)

var nodeTypeToTsType = map[string]string{
	node.FloatType: "number",
	node.StringType: "string",
	node.ListType: "[]",
	node.StructType: "interface",
}

func GenerateTs(n node.Node) string {
	code, extractCode := genTsCode(n)
	if n.Type == node.ListType && strings.Contains(code, "[]") {
		code = fmt.Sprintf("type %s = %s", n.Name, code)
	} else {
		code = ""
	}
	code = fmt.Sprintf("%s\n%s", extractCode, code)
	return strings.Trim(code, "\n")
}

func genTsCode(n node.Node) (string, string) {
	var code string
	var extractCode string

	switch n.Type {
	case node.StringType:
		code = nodeTypeToTsType[node.StringType]
		break
	case node.FloatType:
		code = nodeTypeToTsType[node.FloatType]
		break
	case node.ListType:
		children := n.Children.([]node.Node)
		var childrenType []string
		for _, child := range children {
			childCode, extractChildType := genTsCode(child)
			extractCode += fmt.Sprintf("%s\n", extractChildType)
			childrenType = append(childrenType, childCode)
		}
		// normalize 数组类型
		listType, isInline := normalizeListTypes(childrenType)
		if isInline {
			code = listType
		} else {
			extractCode += fmt.Sprintf("type %s = %s\n", n.Name, listType)
			code = n.Name
		}
		break
	case node.StructType:
		children := n.Children.(map[string]node.Node)
		childrenType := map[string]string{}
		for k, child := range children {
			childCode, extractChildType := genTsCode(child)
			childrenType[k] = childCode
			extractCode += extractChildType
		}
		var structType = fmt.Sprintf("%s %s {\n",nodeTypeToTsType[node.StructType], n.Name)
		for k, childType := range childrenType {
			structType += fmt.Sprintf("    %s: %s\n", k, childType)
		}
		structType += "}"
		extractCode += structType
		code = n.Name
		break
	}
	return code, extractCode
}

func normalizeListTypes(types []string) (string, bool) {
	if len(types) == 0 {
		return "any[]", true
	}
	// 检测数组类型是否一致
	if checkListAllTypes(types) {
		return fmt.Sprintf("%s[]", types[0]), true
	}
	listTypeStr := strings.Join(types, ", ")
	return fmt.Sprintf("[%s]", listTypeStr), false
}

func checkListAllTypes(types [] string) bool {
	var flagType string
	for _, t := range types {
		if flagType == "" {
			flagType = t
			continue
		}
		if t != flagType {
			return false
		}
	}
	return true
}