package abstract_factory

//定义装扮
type Dress struct {
	Cloth ICloth
	Bag IBag
}
//定义装扮工厂
type DressFactory interface {
	GetCloth() ICloth
	GetBag() IBag
}
//定义运动装扮工厂
type SportDressFactory struct {

}

func (sdf SportDressFactory) GetCloth() ICloth {
	return SportCloth{}
}

func (sdf SportDressFactory) GetBag() IBag {
	return SportBag{}
}
//定义休闲装扮工厂
type CasualDressFactory struct {

}

func (cdf CasualDressFactory) GetCloth() ICloth {
	return CasualCloth{}
}

func (cdf CasualDressFactory) GetBag() IBag {
	return CasualBag{}
}






