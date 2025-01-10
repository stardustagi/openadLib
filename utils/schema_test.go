package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Point struct {
	X int
	Y int
	Z int
}

func TestGetFields(t *testing.T) {
	// test code
	p := Point{1, 2, 3}
	fields := getFields(reflect.TypeOf(p))
	for _, field := range fields {
		r, e := json.Marshal(field)
		if e != nil {
			t.Error(e)
		}
		t.Log(string(r))
	}
	tps := getType(reflect.TypeOf(p))
	t.Log(tps)
}
