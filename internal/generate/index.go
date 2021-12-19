package generate

import (
	"fmt"
	"github.com/hubvue/json2type/internal/generate/common"
	"github.com/hubvue/json2type/internal/generate/generator"
	"github.com/hubvue/json2type/internal/node"
	"github.com/hubvue/json2type/internal/util"
	"strings"
)

// Generate the type of the corresponding language through the type node tree
func Generate(n node.Node, langGen generator.Generator) string {
	extractCodeMap := map[string]common.ExtractCode{}
	code := genCode(n, extractCodeMap, langGen)
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
func genCode(n node.Node, extractCodeMap map[string]common.ExtractCode, langGen generator.Generator) string {
	var code string
	switch n.Type {
	case node.StringType:
		code = langGen.GenStringType(n, extractCodeMap)
		break
	case node.FloatType:
		code = langGen.GenFloatType(n, extractCodeMap)
		break
	case node.BoolType:
		code = langGen.GenBoolType(n, extractCodeMap)
		break
	case node.ListType:
		children := n.Children.([]node.Node)
		var childrenType []string
		for _, child := range children {
			childCode := genCode(child, extractCodeMap, langGen)
			childrenType = append(childrenType, childCode)
		}
		code = langGen.GenListType(n, childrenType, extractCodeMap)
		break
	case node.StructType:
		children := n.Children.(map[string]node.Node)
		childrenType := map[string]string{}
		for k, child := range children {
			childCode := genCode(child, extractCodeMap, langGen)
			childrenType[k] = childCode
		}
		code = langGen.GenStructType(n, childrenType, extractCodeMap)
		break
	}
	return code
}
