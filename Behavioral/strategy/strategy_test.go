package strategy

import "testing"

func TestNewStrategy(t *testing.T) {
	addCon := NewContext(&OperationAdd{})
	subCon := NewContext(&OperationSub{})
	mulCon := NewContext(&OperationMultiply{})

	if addCon.ExecStrategy(1,2) != 3{
		t.Error("add Startegy error")
	}

	if subCon.ExecStrategy(1,2) != -1{
		t.Error("sub Startegy error")
	}

	if mulCon.ExecStrategy(1,2) != 2{
		t.Error("multiply Startegy error")
	}
}
