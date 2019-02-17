package factory_method

//定义奔驰车
type BenzCar struct {

}

func (benz BenzCar) GetCarBrand() string{
	return "Benz"
}

//定义奔驰车工厂
type BenzFactory struct {

}

func (bf BenzFactory) CreateCar() ICar {
	return BenzCar{}
}

