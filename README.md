# json2type
convert JSON of a specific format to a type structure(Typescript type or more)


## Quick Start

### CLI

#### Install

> go install github.com/hubvue/json2type/json2type@latest

#### Usage

```txt
Usage of json2type:
  -input string
    	the file of the json file
  -name string
    	the name of the type name (default "auto")
  -output string
    	the name of the file to write the output to (outputs to output.[ext] by default) (default "output")
```

> json2type -input=tmp.json

##### Example
```json
{
  "name": "hubvue",
  "age": 18
}
```
Result:
```ts
interface Auto {
    name: string
    age: number
}
```

### Package

> go get github.com/hubvue/json2type

#### Usage

```go
package main
import (
	"fmt"
	"github.com/hubvue/json2type"
)

func main() {
	fileJson, _ := ioutil.ReadFile(jsonFile)
	// currently only typescript is supported
	code, _ := json2type.Parser(fileJson, "typescript", "typeName")
	fmt.Println(code)
}
```



