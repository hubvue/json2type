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
	code, err := json2type.Parser(fileJson, "typescript", "auto")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)
	//jsonStr, err := json.Marshal([]string{"string", "number", "Text"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(jsonStr))
	//var list []string
	//fmt.Println(list)
}
