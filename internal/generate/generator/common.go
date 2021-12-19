package generator

import (
	"github.com/hubvue/json2type/internal/generate/common"
	"github.com/hubvue/json2type/internal/node"
)

type Generator interface {
	GenStringType(n node.Node, extractCodeMap map[string]common.ExtractCode) string
	GenFloatType(n node.Node, extractCodeMap map[string]common.ExtractCode) string
	GenBoolType(n node.Node, extractCodeMap map[string]common.ExtractCode) string
	GenListType(n node.Node, childrenTypeName []string, extractCodeMap map[string]common.ExtractCode) string
	GenStructType(n node.Node, childrenTypeMap map[string]string, extractCodeMap map[string]common.ExtractCode) string
}
