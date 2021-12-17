package json2type

import (
	"encoding/json"
	"errors"
	"github.com/hubvue/json2type/generator"
	"github.com/hubvue/json2type/node"
)

// Parser resolve JSON to the type of the corresponding language
//	input: JSON data
//	language: type lang
// 	name: outermost layer typeName
func Parser(input []byte, language string, name string) (string, error) {
	var decodeResult interface{}
	err := json.Unmarshal(input, &decodeResult)
	if err != nil {
		return "", errors.New("json decoder error: " + err.Error())
	}
	node := node.Structure2Node(name, decodeResult)

	result := generator.GenerateCode(node, language)
	return result, nil
}
