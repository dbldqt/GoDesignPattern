package simple_factory

//定义奔驰车
type BenzCar struct {

}

func (benz BenzCar) GetCarBrand() string{
	return "Benz"
}
