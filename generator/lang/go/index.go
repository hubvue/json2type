package golang

import (
	"fmt"
	"github.com/hubvue/json2type/generator/common"
	"github.com/hubvue/json2type/node"
	"github.com/hubvue/json2type/util"
	"sort"
	"strings"
)

var NodeTypeToType = map[string]string{
	node.FloatType:  "float64",
	node.StringType: "string",
	node.BoolType:   "bool",
	node.ListType:   "[]interface{}",
	node.StructType: "struct",
}

// ExtractListType get list type-name and extract common type
func ExtractListType(n node.Node, childrenTypeName []string, extractCodeMap map[string]common.ExtractCode) string {
	var childrenExtractType [][]string
	// get extract type children type
	for _, typeName := range childrenTypeName {
		if isInlineType(typeName) {
			return "[]interface{}"
		}
		extractKey, ok := getExtractKeyByName(typeName, extractCodeMap)
		if ok {
			itemTypes := strings.Split(extractKey, "|")
			for _, item := range itemTypes {
				childrenExtractType = append(childrenExtractType, strings.Split(item, "-"))
			}
			delete(extractCodeMap, extractKey)
		}
	}

	childrenTypeMap := make(map[string]struct{})
	var extractItemKey []string
	extractTypeName := util.SnakeToCamel(n.Name, true) + "Item"
	code := fmt.Sprintf("type %s struct {\n", extractTypeName)
	for _, extractType := range childrenExtractType {
		if _, ok := childrenTypeMap[extractType[0]]; !ok {
			childrenTypeMap[extractType[0]] = struct{}{}
			// extractType[0]: Name; extractType[1]: Type
			extractItemKey = append(extractItemKey, fmt.Sprintf("%s-%s", extractType[0], extractType[1]))
			code += fmt.Sprintf("    %s %s `json:\"%s,omitempty\"`\n", util.SnakeToCamel(extractType[0], true), extractType[1], extractType[0])
		}
	}
	extractKey := strings.Join(extractItemKey, "|")
	code += "}\n"
	extractCodeMap[extractKey] = common.ExtractCode{
		Name: extractTypeName,
		Code: code,
	}

	return "[]" + extractTypeName
}

// ExtractStructType get struct type-name and extract common type
func ExtractStructType(n node.Node, childrenTypeMap map[string]string, extractCodeMap map[string]common.ExtractCode) (typeName string) {
	var extractList []string
	for childName, childType := range childrenTypeMap {
		extractList = append(extractList, fmt.Sprintf("%s-%s", childName, childType))
	}
	sort.Strings(extractList)
	extractKey := strings.Join(extractList, "|")
	extractValue, ok := extractCodeMap[extractKey]
	if ok {
		typeName = extractValue.Name
	} else {
		typeName = util.SnakeToCamel(n.Name, true)
		var typeCode = fmt.Sprintf("type %s struct {\n", typeName)
		for k, childType := range childrenTypeMap {
			typeCode += fmt.Sprintf("    %s %s `json:\"%s\"`\n", util.SnakeToCamel(k, true), childType, k)
		}
		typeCode += "}\n"
		extractCodeMap[extractKey] = common.ExtractCode{
			Name: typeName,
			Code: typeCode,
		}
	}
	return typeName
}

// getExtractKeyByName  get extract key by name
func getExtractKeyByName(typeName string, extractCodeMap map[string]common.ExtractCode) (string, bool) {
	for k, extractCode := range extractCodeMap {
		if extractCode.Name == typeName {
			return k, true
		}
	}
	return "", false
}

// isInlineType base type and list type
func isInlineType(typeName string) bool {
	if typeName == NodeTypeToType[node.BoolType] ||
		typeName == NodeTypeToType[node.FloatType] ||
		typeName == NodeTypeToType[node.StringType] ||
		isInlineListType(typeName) {
		return true
	}
	return false
}

// isInlineListType list type
func isInlineListType(typeName string) bool {
	if strings.Contains(typeName, "[]") {
		return true
	}
	return false
}
