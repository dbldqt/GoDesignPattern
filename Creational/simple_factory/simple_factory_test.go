package simple_factory

import "testing"

func TestNewClass(t *testing.T) {
	if NewClass("Bmw").GetCarBrand() != "Bmw" || NewClass("Benz").GetCarBrand() != "Benz"{
		t.Error("faactory fail")
	}
}