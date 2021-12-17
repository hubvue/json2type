package generator

import (
	"fmt"
	"github.com/hubvue/json2type/generator/common"
	golang "github.com/hubvue/json2type/generator/lang/go"
	"github.com/hubvue/json2type/generator/lang/ts"
	"github.com/hubvue/json2type/node"
	"github.com/hubvue/json2type/util"
	"strings"
)

var langType = map[string]map[string]string{
	"typescript": ts.NodeTypeToType,
	"go":         golang.NodeTypeToType,
}

// GenerateCode generate the type of the corresponding language through the type node tree
func GenerateCode(n node.Node, lang string) string {
	extractCodeMap := map[string]common.ExtractCode{}
	code := genCode(n, extractCodeMap, lang)
	if n.Type == node.ListType && strings.Contains(code, "[]") {
		code = fmt.Sprintf("type %s = %s", util.SnakeToCamel(n.Name, true), code)
	} else {
		code = ""
	}
	extractTypeCode := common.GenExtractTypeCode(extractCodeMap)
	code = fmt.Sprintf("%s\n%s", extractTypeCode, code)
	return strings.Trim(code, "\n")
}

// genCode traversing the type node tree to generate types
func genCode(n node.Node, extractCodeMap map[string]common.ExtractCode, lang string) string {
	var code string
	switch n.Type {
	case node.StringType:
		code = langType[lang][node.StringType]
		break
	case node.FloatType:
		code = langType[lang][node.FloatType]
		break
	case node.BoolType:
		code = langType[lang][node.BoolType]
		break
	case node.ListType:
		children := n.Children.([]node.Node)
		var childrenType []string
		for _, child := range children {
			childCode := genCode(child, extractCodeMap, lang)
			childrenType = append(childrenType, childCode)
		}
		code = extractType(n, childrenType, extractCodeMap, lang)
		break
	case node.StructType:
		children := n.Children.(map[string]node.Node)
		childrenType := map[string]string{}
		for k, child := range children {
			childCode := genCode(child, extractCodeMap, lang)
			childrenType[k] = childCode
		}
		code = extractType(n, childrenType, extractCodeMap, lang)
		break
	}
	return code
}

// extractType get type-name and extract common type
func extractType(n node.Node, childrenType interface{}, extractCodeMap map[string]common.ExtractCode, lang string) (typeName string) {
	switch n.Type {
	case node.ListType:
		childrenTypeName := childrenType.([]string)
		switch lang {
		case "go":
			typeName = golang.ExtractListType(n, childrenTypeName, extractCodeMap)
		case "typescript":
			typeName = ts.ExtractListType(n, childrenTypeName, extractCodeMap)
		}
	case node.StructType:
		childrenTypeMap := childrenType.(map[string]string)
		switch lang {
		case "go":
			typeName = golang.ExtractStructType(n, childrenTypeMap, extractCodeMap)
		case "typescript":
			typeName = ts.ExtractStructType(n, childrenTypeMap, extractCodeMap)
		}

	}
	return typeName
}
