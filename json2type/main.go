package main

import (
	"flag"
	"fmt"
	"github.com/hubvue/json2type"
	"io/ioutil"
	"os"
	"path"
)

var (
	input    = flag.String("input", "", "the file of the json file(input parameter is required)")
	name     = flag.String("name", "auto", "the name of the type name(auto by default)")
	output   = flag.String("output", "output", "the name of the file to write the output to (outputs to output.[ext] by default)")
	language = flag.String("language", "typescript", "used to convert json to the type of the language(typescript by default)")
)

var extMap = map[string]string{
	"go":         "go",
	"typescript": "ts",
}

func main() {
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		errorHandler("failed to get current directory", nil)
	}
	filePath := pwd + "/" + *input
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) || path.Ext(filePath) != ".json" {
		flag.Usage()
		errorHandler("file not found or is not json file", nil)
	}

	fileJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		errorHandler("read json file err: ", err)
	}
	code, err := json2type.Parser(fileJson, *language, *name)
	if err != nil {
		errorHandler("parser json err: ", err)
	}
	err = ioutil.WriteFile(*output+"."+extMap[*language], []byte(code), 0777)
	if err != nil {
		errorHandler("output file err: ", err)
	}
}

func errorHandler(message string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, message, err)
	}
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
