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
	err = ioutil.WriteFile("tmp.ts", []byte(code), 0777)
	if err != nil {
		fmt.Println("write file error: ", err)
	}
}
