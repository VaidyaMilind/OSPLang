package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("string with same content have diiferent hash key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("String with same content have diff hash key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("String with differnent content havev same hash key")
	}
}
