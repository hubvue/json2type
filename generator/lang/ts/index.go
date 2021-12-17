package ts

import (
	"fmt"
	"github.com/hubvue/json2type/generator/common"
	"github.com/hubvue/json2type/node"
	"github.com/hubvue/json2type/util"
	"sort"
	"strings"
)

// NodeTypeToType node type to ts type
var NodeTypeToType = map[string]string{
	node.FloatType:  "number",
	node.StringType: "string",
	node.BoolType:   "boolean",
	node.ListType:   "[]",
	node.StructType: "interface",
}

// ExtractListType get list type-name and extract common type
func ExtractListType(n node.Node, childrenTypeName []string, extractCodeMap map[string]common.ExtractCode) (typeName string) {
	listType, isInline := normalizeListTypes(childrenTypeName)
	if isInline {
		typeName = listType
	} else {
		extractKey := strings.Join(childrenTypeName, "-")
		extractValue, ok := extractCodeMap[extractKey]
		if ok {
			typeName = extractValue.Name
		} else {
			name := util.SnakeToCamel(n.Name, true)
			typeName = name
			typeCode := fmt.Sprintf("type %s = %s\n", name, listType)
			extractCodeMap[extractKey] = common.ExtractCode{
				Name: name,
				Code: typeCode,
			}
		}
	}
	return typeName
}

// normalizeListTypes normalize list type
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

// checkListAllTypes check for consistency of list subtypes
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

// ExtractStructType get struct type-name and extract common type
func ExtractStructType(n node.Node, childrenTypeMap map[string]string, extractCodeMap map[string]common.ExtractCode) (typeName string) {
	var extractList []string
	for childName, childType := range childrenTypeMap {
		extractList = append(extractList, fmt.Sprintf("%s-%s", childName, childType))
	}
	// fix: solving the map random traversal problem
	sort.Strings(extractList)
	extractKey := strings.Join(extractList, "|")
	extractValue, ok := extractCodeMap[extractKey]
	if ok {
		typeName = extractValue.Name
	} else {
		typeName = util.SnakeToCamel(n.Name, true)
		var typeCode = fmt.Sprintf("%s %s {\n", NodeTypeToType[node.StructType], typeName)
		for k, childType := range childrenTypeMap {
			typeCode += fmt.Sprintf("    %s: %s\n", util.SnakeToCamel(k, false), childType)
		}
		typeCode += "}\n"
		extractCodeMap[extractKey] = common.ExtractCode{
			Name: typeName,
			Code: typeCode,
		}
	}

	return typeName
}
