package json2type

import (
	"errors"
	"fmt"
	"github.com/hubvue/json2type/internal/generate"
	"github.com/hubvue/json2type/internal/generate/generator"
	"github.com/hubvue/json2type/internal/parser"
)

// Run resolve JSON to the type of the corresponding language
//	input: JSON data
//	language: type lang
// 	name: outermost layer typeName
func Run(input []byte, language string, name string) (string, error) {
	node, err := parser.Parser(input, name)
	if err != nil {
		return "", err
	}
	langGen := getLangGenerator(language)
	if langGen == nil {
		return "", errors.New(fmt.Sprintf("not found %s generator", language))
	}
	result := generate.Generate(node, langGen)
	return result, nil
}

func getLangGenerator(lang string) generator.Generator {
	var langGen generator.Generator
	switch lang {
	case "typescript":
		langGen = generator.NewTsGenerator()
	case "go":
		langGen = generator.NewGoGenerator()
	}
	return langGen
}
