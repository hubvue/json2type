# json2type
convert JSON of a specific format to a type structure(support Typescript type or Go type)


## Quick Start

### CLI

#### Install

use go tool install

```shell
go install github.com/hubvue/json2type/json2type@latest
```

use npm/yarn/pnpm install
```shell
# npm
npm install @cckim/json2type -g
# yarn
yarn global add @cckim/json2type
# pnpm
pnpm add @cckim/json2type --global
```
#### Usage

```shell
json2type help
```

```txt
Usage of json2type:
  -input string
    	the file of the json file
  -name string
    	the name of the type name (default "auto")
  -output string
    	the name of the file to write the output to (outputs to output.[ext] by default) (default "output")
```

```shell
json2type -input=tmp.json
```

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

```shell
go get github.com/hubvue/json2type
```
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



