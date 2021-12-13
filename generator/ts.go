package generator

import (
	"fmt"
	"github.com/hubvue/json2type/node"
	"github.com/hubvue/json2type/util"
	"sort"
	"strings"
)

type extractCode struct {
	Name string
	Code string
}

var nodeTypeToTsType = map[string]string{
	node.FloatType:  "number",
	node.StringType: "string",
	node.BoolType:   "boolean",
	node.ListType:   "[]",
	node.StructType: "interface",
}

func GenerateTs(n node.Node) string {
	extractCodeMap := map[string]extractCode{}
	code := genTsCode(n, extractCodeMap)
	if n.Type == node.ListType && strings.Contains(code, "[]") {
		code = fmt.Sprintf("type %s = %s", util.SnakeToCamel(n.Name, true), code)
	} else {
		code = ""
	}
	extractTypeCode := genExtractTypeCode(extractCodeMap)
	code = fmt.Sprintf("%s\n%s", extractTypeCode, code)
	return strings.Trim(code, "\n")
}

func genTsCode(n node.Node, extractCodeMap map[string]extractCode) string {
	var code string
	switch n.Type {
	case node.StringType:
		code = nodeTypeToTsType[node.StringType]
		break
	case node.FloatType:
		code = nodeTypeToTsType[node.FloatType]
		break
	case node.BoolType:
		code = nodeTypeToTsType[node.BoolType]
		break
	case node.ListType:
		children := n.Children.([]node.Node)
		var childrenType []string
		for _, child := range children {
			childCode := genTsCode(child, extractCodeMap)
			childrenType = append(childrenType, childCode)
		}
		code = extractType(n, childrenType, extractCodeMap)
		break
	case node.StructType:
		children := n.Children.(map[string]node.Node)
		childrenType := map[string]string{}
		for k, child := range children {
			childCode := genTsCode(child, extractCodeMap)
			childrenType[k] = childCode
		}
		code = extractType(n, childrenType, extractCodeMap)
		break
	}
	return code
}

func genExtractTypeCode(extractCodeMap map[string]extractCode) (code string) {
	for _, extractValue := range extractCodeMap {
		code += extractValue.Code
	}
	return code
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

func checkListAllTypes(types []string) bool {
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

func extractType(n node.Node, childrenType interface{}, extractCodeMap map[string]extractCode) (typeName string) {
	switch n.Type {
	case node.ListType:
		types := childrenType.([]string)
		listType, isInline := normalizeListTypes(types)
		if isInline {
			typeName = listType
		} else {
			extractKey := strings.Join(types, "-")
			extractValue, ok := extractCodeMap[extractKey]
			if ok {
				typeName = extractValue.Name
			} else {
				name := util.SnakeToCamel(n.Name, true)
				typeName = name
				typeCode := fmt.Sprintf("type %s = %s\n", name, listType)
				extractCodeMap[extractKey] = extractCode{
					Name: name,
					Code: typeCode,
				}
			}
		}
	case node.StructType:
		types := childrenType.(map[string]string)
		var extractList []string
		for childName, childType := range types {
			extractList = append(extractList, fmt.Sprintf("%s-%s", childName, childType))
		}
		// fix: solving the map random traversal problem
		sort.Strings(extractList)
		extractKey := strings.Join(extractList, "-")
		extractValue, ok := extractCodeMap[extractKey]
		if ok {
			typeName = extractValue.Name
		} else {
			name := util.SnakeToCamel(n.Name, true)
			typeName = name
			var typeCode = fmt.Sprintf("%s %s {\n", nodeTypeToTsType[node.StructType], name)
			for k, childType := range types {
				typeCode += fmt.Sprintf("    %s: %s\n", util.SnakeToCamel(k, false), childType)
			}
			typeCode += "}\n"
			extractCodeMap[extractKey] = extractCode{
				Name: name,
				Code: typeCode,
			}
		}
	}
	return typeName
}