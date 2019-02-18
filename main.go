package main

import "GoDesignPattern/Creational/builder"

func main(){
	director := builder.Director{
		CarBuilder:builder.CarBuilder{},
	}

	//给我一辆奔驰suv
	benzSuv := director.CreateBenzSuv()
	if benzSuv.GetWheel() != "bbenz的轮胎" || benzSuv.GetEngine() != "benz的引擎"{
		println("benzSuv error")
	}

	//给我一辆宝马商务车
	bmwVan := director.CreateBWMVan()
	if bmwVan.GetWheel() != "BMW的轮胎" || bmwVan.GetEngine() != "BMW的引擎"{
		println("benzSuv error")
	}

	//给我一辆混合车型
	complexCar := director.CreateComplexCar()
	if complexCar.GetWheel() != "benz的轮胎" || complexCar.GetEngine() != "BMW的引擎"{
		println("benzSuv error")
	}
}

