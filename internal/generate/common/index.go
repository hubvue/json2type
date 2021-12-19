package common

type ExtractCode struct {
	Name string
	Code string
}

// GenExtractTypeCode generate extract type code
func GenExtractTypeCode(extractCodeMap map[string]ExtractCode) (code string) {
	for _, extractValue := range extractCodeMap {
		code += extractValue.Code
	}
	return code
}
