package abstract_factory

/**
	本例假定一个少年出去聚会，需要搭配同种风格的背包和衣服
 */
//基础接口，其本身及其方法均为导出的方法，工厂所产生的所有子类对象都实现该接口
func NewDress(dressFactory DressFactory) Dress{
	return Dress{
		Cloth:dressFactory.GetCloth(),
		Bag:dressFactory.GetBag(),
	}
}












