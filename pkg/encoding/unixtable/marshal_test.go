package unixtable_test

import (
	"strings"
	"testing"

	"github.com/frantjc/forge/pkg/encoding/unixtable"
)

func TestMarshal(t *testing.T) {
	var (
		ut = &Unixtable{
			One: "hello",
			Two: "there",
		}
		expected = []byte(`One     Two
hello   there
`)
	)

	actual, err := unixtable.Marshal(ut)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var (
		actualStr   = string(actual)
		expectedStr = string(expected)
	)

	if !strings.EqualFold(actualStr, expectedStr) {
		t.Error("actual\n", actualStr, "does not equal expected\n", expectedStr)
		t.FailNow()
	}
}

func TestMarshalTag(t *testing.T) {
	var (
		ut = &UnixtableTagged{
			One: "hello",
			Two: "there",
		}
		expected = []byte(`one     two
hello   there
`)
	)

	actual, err := unixtable.Marshal(ut)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	var (
		actualStr   = string(actual)
		expectedStr = string(expected)
	)

	if !strings.EqualFold(actualStr, expectedStr) {
		t.Error("actual\n", actualStr, "does not equal expected\n", expectedStr)
		t.FailNow()
	}
}
