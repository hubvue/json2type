package generator

import (
	"github.com/hubvue/json2type/internal/generate/common"
	"github.com/hubvue/json2type/internal/node"
)

var nodeTypeToRustType = map[string]string{
	node.FloatType:  "f64",
	node.StringType: "String",
	node.BoolType:   "bool",
	node.ListType:   "()",
	node.StructType: "pub struct",
}

type rustGenerator struct{}

func NewRustGenerator() *rustGenerator {
	return &rustGenerator{}
}

func (gen *rustGenerator) GenFloatType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToRustType[node.FloatType]
}

func (gen *rustGenerator) GenStringType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToRustType[node.StringType]
}

func (gen *rustGenerator) GenBoolType(n node.Node, extractCodeMap map[string]common.ExtractCode) string {
	return nodeTypeToRustType[node.BoolType]
}
