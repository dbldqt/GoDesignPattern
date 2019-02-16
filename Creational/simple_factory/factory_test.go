package simple_factory

import "testing"

func TestNewClass(t *testing.T) {
	if NewClass("A").GetClassName() != "ClassA" || NewClass("B").GetClassName() != "ClassB"{
		t.Error("faactory fail")
	}
}