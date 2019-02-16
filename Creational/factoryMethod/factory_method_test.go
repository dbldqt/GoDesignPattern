package simpleFatory

import "testing"

func TestNewClass(t *testing.T) {
	if (&FactoryC{}).GetInstance().GetClassName() != "classC"{
		t.Error("classC fail")
	}

	if (&FactoryB{}).GetInstance().GetClassName() != "classB"{
		t.Error("classB fail")
	}

	if (&FactoryA{}).GetInstance().GetClassName() != "classA"{
		t.Error("classA fail")
	}
}