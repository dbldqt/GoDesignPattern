package factory_method

import "testing"

func TestNewClass(t *testing.T) {
	bmwF := BmwFactory{}
	if bmwF.CreateCar().GetCarBrand() != "Bmw" {
		t.Error("Bmw factory fail")
	}

	benzF := BenzFactory{}
	if benzF.CreateCar().GetCarBrand() != "Benz"{
		t.Error("Benz factory fail")
	}

	//如果需要其他品牌的车，则只需定义新的汽车品牌和汽车工厂即可，无需改动现有代码
}