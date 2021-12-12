package main

import (
	"fmt"
	"github.com/hubvue/json2type"
	"io/ioutil"
)

func main() {
	fileJson, err := ioutil.ReadFile("../json/list.json")
	if err != nil {
		fmt.Println(err)
	}
	json2type.Parser(fileJson, "go", "auto")
}
