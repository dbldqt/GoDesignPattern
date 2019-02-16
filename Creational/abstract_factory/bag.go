package abstract_factory

type IBag interface {
	GetBagStyle() string
}

//定义运动风格背包
type SportBag struct {

}
func (sb SportBag) GetBagStyle() string {
	return "sportBag"
}
//定义休闲风格背包
type CasualBag struct {

}
func (cb CasualBag) GetBagStyle() string {
	return "casualBag"
}

////定义背包生成工厂
//type BagFactory interface {
//	GetBag() IBag
//}
//
////定义生成运动风格衣服的生成工厂
//type sportBagFactory struct {
//
//}
//
//func (scf sportBagFactory) GetBag() IBag {
//	return SportBag{}
//}
////定义生成休闲风格的衣服工厂
//type casualBagFactory struct {
//
//}
//
//func (ccf casualBagFactory) GetBag() IBag {
//	return CasualBag{}
//}


