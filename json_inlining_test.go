package goissues

import (
	"encoding/json"
	"testing"
)

func TestJSONInlining(t *testing.T) {
	type Bar struct {
		Baz string
	}
	type Foo struct {
		*Bar `json:",omitempty"`
	}
	foo := &Foo{
		Bar: &Bar{
			Baz: "baz",
		},
	}
	b, err := json.Marshal(foo)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"Baz":"baz"}` {
		t.Fatalf("Got %s", b)
	}
}
