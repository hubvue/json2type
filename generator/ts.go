package generator

import "github.com/hubvue/json2type/node"

func GenerateTs(n node.Node) string {
	typesMap := map[string]string{}
	code, _ := genTsCode(n, &typesMap)
	return code
}

func genTsCode(n node.Node, typeMap *map[string]string) (code string, typeName string) {
	//switch
	return "", ""
}