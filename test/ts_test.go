package test

import (
	"github.com/hubvue/json2type"
	"strings"
	"testing"
)

func TestJson2Ts(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, got string) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %s, but got %s", expected, got)
		}
	}

	t.Run("parser simple list", func(t *testing.T) {
		json := []byte(`["string", "number", "Text"]`)
		expected := `type Auto = string[]`
		code, _ := json2type.Run(json, "typescript", "auto")
		assertCorrectMessage(t, expected, code)
	})
	t.Run("parser complex list", func(t *testing.T) {
		json := []byte(`["string", 1, true]`)
		expected := `type Auto = [string, number, boolean]`
		code, _ := json2type.Run(json, "typescript", "auto")
		assertCorrectMessage(t, expected, code)
	})
	t.Run("parser embedded map list", func(t *testing.T) {
		json := []byte(`[{"name": "kim"}, {"age": 123}]`)
		code, _ := json2type.Run(json, "typescript", "auto")
		if !strings.Contains(code, "Child0") {
			t.Error("parser list error: not found Child0")
		}
		if !strings.Contains(code, "Child1") {
			t.Error("parser list error: not found Child1 ")
		}
		if !strings.Contains(code, "type Auto = [AutoChild0, AutoChild1]") {
			t.Error("parser list error: not found Auto type")
		}
	})
	t.Run("parser embedded same map list", func(t *testing.T) {
		json := []byte(`[{"name": "kim"}, {"name": "hubvue"}]`)
		code, _ := json2type.Run(json, "typescript", "auto")
		if !strings.Contains(code, "AutoChild0") {
			t.Error("parser list error: not found AutoChild0")
		}
		if strings.Contains(code, "AutoChild1") {
			t.Error("parser list error: found AutoChild1")
		}
		if !strings.Contains(code, "type Auto = AutoChild0[]") {
			t.Error("parser list error: not found Auto type")
		}
	})
	t.Run("parser simple map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "age": 18}`)
		code, _ := json2type.Run(json, "typescript", "auto")
		if !strings.Contains(code, "interface") {
			t.Error("parser map error: not found interface keyword")
		}
		if !strings.Contains(code, "name") || !strings.Contains(code, "string") {
			t.Error("parser map error: not found name and string type")
		}
		if !strings.Contains(code, "age") || !strings.Contains(code, "number") {
			t.Error("parser map error not found age and number type")
		}
	})
	t.Run("parser embedded simple list map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "keys": ["string", "number"]}`)
		code, _ := json2type.Run(json, "typescript", "auto")
		if !strings.Contains(code, "string[]") {
			t.Error("parser map error: not found string[]")
		}
	})
	t.Run("parser embedded complex list map", func(t *testing.T) {
		json := []byte(`{"name": "kim", "keys": ["string", 123, true]}`)
		code, _ := json2type.Run(json, "typescript", "auto")
		if !strings.Contains(code, "keys: Keys") {
			t.Error("parser map error: not found keys type")
		}
		if !strings.Contains(code, "type Keys = [string, number, boolean]") {
			t.Error("parser map error: not found Keys")
		}
	})
}
