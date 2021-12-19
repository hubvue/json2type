package parser

import (
	"encoding/json"
	"errors"
	"github.com/hubvue/json2type/internal/node"
)

// Parser parsing JSON into specific generic structures
func Parser(input []byte, name string) (node.Node, error) {
	var decodeResult interface{}
	err := json.Unmarshal(input, &decodeResult)
	if err != nil {
		return node.Node{}, errors.New("json decoder error: " + err.Error())
	}
	node := node.Structure2Node(name, decodeResult)
	return node, nil
}
