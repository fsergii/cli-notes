package main

import (
	"reflect"
	"testing"
)

func TestCreateRecord(t *testing.T) {
	want := Record{
		Note: "test",
		id:   1,
	}
	msg, err := createRecord("test")

	if !reflect.DeepEqual(want, msg) || err != nil {
		t.Fatalf(`createRecord("test") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
