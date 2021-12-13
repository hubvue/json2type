package test

import (
	"github.com/hubvue/json2type/util"
	"testing"
)

func TestUtil(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, got string) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %s, but got %s", expected, got)
		}
	}

	t.Run("test SnakeToCamel", func(t *testing.T) {
		t.Run("firstUpper is true", func(t *testing.T) {
			value := "abc_abc"
			expected := "AbcAbc"
			got := util.SnakeToCamel(value, true)
			assertCorrectMessage(t, expected, got)
		})
		t.Run("firstUpper is false", func(t *testing.T) {
			value := "abc_abc"
			expected := "abcAbc"
			got := util.SnakeToCamel(value, false)
			assertCorrectMessage(t, expected, got)
		})
	})
}
