package test

import (
	"github.com/hubvue/json2type"
	"strings"
	"testing"
)

func TestJson2Go(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, got string) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %s, but got %s", expected, got)
		}
	}

	t.Run("parser simple list", func(t *testing.T) {
		json := []byte(`["string", "number", "Text"]`)
		expected := `type Auto = []interface{}`
		code, _ := json2type.Parser(json, "go", "auto")
		assertCorrectMessage(t, expected, code)
	})
	t.Run("parser complex list", func(t *testing.T) {
		json := []byte(`["string", 1, true]`)
		expected := `type Auto = []interface{}`
		code, _ := json2type.Parser(json, "go", "auto")
		assertCorrectMessage(t, expected, code)
	})
	t.Run("parser embedded map list", func(t *testing.T) {
		json := []byte(`[{"name": "kim"}, {"age": 123}]`)
		code, _ := json2type.Parser(json, "go", "auto")
		if !strings.Contains(code, "Name") && !strings.Contains(code, "string") {
			t.Error("parser list error: not found Name and string type")
		}
		if !strings.Contains(code, "Age") && !strings.Contains(code, "float64") {
			t.Error("parser list error: not found Age and float64 type")
		}
		if !strings.Contains(code, "type Auto = []AutoItem") {
			t.Error("parser list error: not found Auto type")
		}
		if !strings.Contains(code, "omitempty") {
			t.Error("parser list error: not found AutoItem key omitempty tag")
		}
	})
	t.Run("parser embedded same map list", func(t *testing.T) {
		json := []byte(`[{"name": "kim"}, {"name": "hubvue"}]`)
		code, _ := json2type.Parser(json, "go", "auto")
		if !strings.Contains(code, "Name") && !strings.Contains(code, "string") {
			t.Error("parser list error: not found Name and string type")
		}
		if !strings.Contains(code, "type Auto = []AutoItem") {
			t.Error("parser list error: not found Auto type")
		}
	})
	t.Run("parser simple map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "age": 18}`)
		code, _ := json2type.Parser(json, "go", "auto")
		if !strings.Contains(code, "struct") {
			t.Error("parser map error: not found struct keyword")
		}
		if !strings.Contains(code, "Name") && !strings.Contains(code, "string") {
			t.Error("parser list error: not found Name and string type")
		}
		if !strings.Contains(code, "Age") && !strings.Contains(code, "float64") {
			t.Error("parser list error: not found Age and float64 type")
		}
		if strings.Contains(code, "omitempty") {
			t.Error("parser list error: found key omitempty tag")
		}
		if !strings.Contains(code, "json") {
			t.Error("parser list error: not found key json tag")
		}
	})
	t.Run("parser embedded simple list map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "keys": ["string", "number"]}`)
		code, _ := json2type.Parser(json, "go", "auto")
		if !strings.Contains(code, "Keys") && !strings.Contains(code, "[]interface{}") {
			t.Error("parser map error: not found Keys and []interface{}")
		}
	})
	t.Run("parser embedded complex list map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "keys": [{ "name": "item1"},{ "age": 123}]}`)
		code, _ := json2type.Parser(json, "go", "auto")
		if !strings.Contains(code, "keys") && !strings.Contains(code, "[]KeysItem") {
			t.Error("parser map error: not found Keys and []KeysItem type")
		}
		if !strings.Contains(code, "KeysItem struct") {
			t.Error("parser map error: not found KeysItem")
		}
	})

}
