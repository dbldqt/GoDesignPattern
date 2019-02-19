package prototype

import "testing"
var manager *PrototypeManager

func init(){
	manager = NewPrototypeManager()
}

func TestClone(t *testing.T){
	ceil1 := Ceil{
		"ceil1",
	}
	manager.set("Ceil",&ceil1)
	suv2 := manager.get("Ceil")

	if ceil1 != ceil1{
		t.Error("self equal error")
	}

	if suv2.(*Ceil) == &ceil1{
		t.Error("prototype manager error")
	}

	suv3 := ceil1.Clone()
	if suv3.(*Ceil) == &ceil1{
		t.Error("clone error")
	}
}
