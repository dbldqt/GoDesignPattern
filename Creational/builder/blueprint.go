package builder

//定义汽车制造蓝图类,因为一种产品只有一个蓝图不需要抽象
/*
	这和一个具体的产品Car类是一样的？错，不一样！它是一个蓝图，是一个可以参考的模板，有一个蓝图可以设计出非常多的产品，
	如有一个R系统的奔驰商务车设计蓝图，我们就可以生产出一系列的奔驰车。它指导我们的产品生产，
	而不是一个具体的产品。我们来看宝马车建造车间。

	blueprint中文的意思是“蓝图”，你要建造一辆车必须有一个设计样稿或者蓝图吧，否则怎么生产？怎么装配？该类就是一个可参考的生产样本。
 */
type CarBluePrint struct {
	wheel string
	engine string
}

func (blue CarBluePrint) GetWheel() string{
	return blue.wheel
}

func (blue CarBluePrint) GetEngine() string{
	return blue.engine
}

func (blue CarBluePrint) SetWheel(wheel string){
	blue.wheel = wheel
}

func (blue CarBluePrint) SetEngine(engine string){
	blue.engine = engine
}





