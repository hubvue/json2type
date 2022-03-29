package generator

import (
	"fmt"
	"go/format"
	"sort"
	"strings"

	"github.com/hubvue/json2type/internal/generate/common"
	"github.com/hubvue/json2type/internal/node"
	"github.com/hubvue/json2type/internal/util"
)

var nodeTypeToGoType = map[string]string{
	node.FloatType:  "float64",
	node.StringType: "string",
	node.BoolType:   "bool",
	node.ListType:   "[]interface{}",
	node.StructType: "pub struct",
}

type goGenerator struct{}

func NewGoGenerator() *goGenerator {
	return &goGenerator{}
}

func (gen *goGenerator) GenStringType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToGoType[node.StringType]
}
func (gen *goGenerator) GenFloatType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToGoType[node.FloatType]
}
func (gen *goGenerator) GenBoolType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToGoType[node.BoolType]
}
func (gen *goGenerator) GenListType(n node.Node, childrenTypeName []string, extractCodeMap map[string]common.ExtractCode) string {
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
	// format
	codeByte, err := format.Source([]byte(code))
	if err == nil {
		code = string(codeByte)
	}
	extractCodeMap[extractKey] = common.ExtractCode{
		Name: extractTypeName,
		Code: code,
	}

	return "[]" + extractTypeName
}
func (gen *goGenerator) GenStructType(n node.Node, childrenTypeMap map[string]string, extractCodeMap map[string]common.ExtractCode) string {
	var typeName string
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
		var code = fmt.Sprintf("type %s struct {\n", typeName)
		for k, childType := range childrenTypeMap {
			code += fmt.Sprintf("    %s %s `json:\"%s\"`\n", util.SnakeToCamel(k, true), childType, k)
		}
		code += "}\n"
		// format
		codeByte, err := format.Source([]byte(code))
		if err == nil {
			code = string(codeByte)
		}
		extractCodeMap[extractKey] = common.ExtractCode{
			Name: typeName,
			Code: code,
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
	if typeName == nodeTypeToGoType[node.BoolType] ||
		typeName == nodeTypeToGoType[node.FloatType] ||
		typeName == nodeTypeToGoType[node.StringType] ||
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
