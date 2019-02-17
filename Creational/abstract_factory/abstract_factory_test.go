package abstract_factory

import "testing"

func TestBenzFactory(t *testing.T) {
	bmwF := BwmFatcory{}
	if bmwF.CreateMpv().GetModel() != "BmwMpv" || bmwF.CreateSuv().GetModel() != "BmwSuv"{
		t.Error("BMW factory error")
	}

	benzF := BenzFactory{}
	if benzF.CreateSuv().GetModel() != "BenzSuv" || benzF.CreateMpv().GetModel() != "BenzMpv"{
		t.Error("Benz factory error")
	}
}