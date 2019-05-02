package qbuilder

import (
	"reflect"
	"testing"

	"github.com/spacycoder/test/cosmos"
)

func TestQueryBuilder(t *testing.T) {
	qb := New()

	params := []cosmos.P{{"@SOMETHING", 20}, {"@NAME", "Lars"}}
	res := qb.Select("*").From("root").And("root.age > @SOMETHING").And("root.name = @NAME").Params(params...).Build()
	if res.Query != "SELECT * FROM root WHERE root.age > @SOMETHING AND root.name = @NAME" {
		t.Errorf("Invalid query string result: %s, should be: %s", res.Query, "SELECT * FROM root WHERE root.age > @SOMETHING AND root.name = @NAME")
	}

	if len(res.Parameters) != len(params) {
		t.Errorf("Invalid count: %d, should be: %d", len(res.Parameters), len(params))
	}

	for i, p := range res.Parameters {
		if p.Name != params[i].Name {
			t.Errorf("Names do not match: %s vs %s", p.Name, params[i].Name)
		}
		if !reflect.DeepEqual(p.Value, params[i].Value) {
			t.Errorf("Values are not equal")
		}
	}

	qb = New()
	params = []cosmos.P{{"@SOMETHING", 20}, {"@NAME", "Lars"}}
	res = qb.Select("*").From("root").And("root.age > @SOMETHING").Or("root.name = @NAME").Params(params...).Build()
	if res.Query != "SELECT * FROM root WHERE root.age > @SOMETHING OR root.name = @NAME" {
		t.Errorf("Invalid query string result: %s, should be: %s", res.Query, "SELECT * FROM root WHERE root.age > @SOMETHING OR root.name = @NAME")
	}

	if len(res.Parameters) != len(params) {
		t.Errorf("Invalid count: %d, should be: %d", len(res.Parameters), len(params))
	}

	for i, p := range res.Parameters {
		if p.Name != params[i].Name {
			t.Errorf("Names do not match: %s vs %s", p.Name, params[i].Name)
		}
		if !reflect.DeepEqual(p.Value, params[i].Value) {
			t.Errorf("Values are not equal")
		}
	}
}
