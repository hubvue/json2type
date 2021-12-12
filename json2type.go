package json2type

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hubvue/json2type/generator"
	"github.com/hubvue/json2type/node"
)

// Parser xxx
func Parser(input []byte, language string, name string) (string, error) {
	var decodeResult interface{}
	err := json.Unmarshal(input, &decodeResult)
	if err != nil {
		return "", errors.New("json decoder error: " + err.Error())
	}
	node := node.Structure2Node(name, decodeResult)

	fmt.Println("node", node)

	var result string
	switch language {
	case "go":
		result = generator.GenerateGo(node)
	case "typescript":
		result = generator.GenerateTs(node)
	}

	return result, nil
}
