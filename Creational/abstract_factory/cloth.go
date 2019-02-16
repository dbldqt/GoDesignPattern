package abstract_factory

type ICloth interface {
	GetClothStyle() string
}

//定义运动风格衣服
type SportCloth struct {

}
func (sc SportCloth) GetClothStyle() string {
	return "sportCloth"
}
//定义休闲风格衣服
type CasualCloth struct {

}
func (cc CasualCloth) GetClothStyle() string {
	return "casualCloth"
}

////定义衣服生成工厂
//type ClothFactory interface {
//	GetCloth() ICloth
//}
////定义生成运动风格衣服的生成工厂
//type sportClothFactory struct {
//
//}
//
//func (scf sportClothFactory) GetCloth() ICloth {
//	return SportCloth{}
//}
////定义生成休闲风格的衣服工厂
//type casualClothFactory struct {
//
//}
//
//func (ccf casualClothFactory) GetCloth() ICloth {
//	return CasualCloth{}
//}