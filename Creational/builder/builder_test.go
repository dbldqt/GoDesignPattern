package builder

import "testing"

func TestCarBuilder_BuildCar(t *testing.T) {
	director := Director{
		carBuilder:CarBuilder{},
	}

	//给我一辆奔驰suv
	benzSuv := director.CreateBenzSuv()
	if benzSuv.GetWheel() != "bbenz的轮胎" || benzSuv.GetEngine() != "benz的引擎"{
		t.Error("benzSuv error")
	}

	//给我一辆宝马商务车
	bmwVan := director.CreateBWMVan()
	if bmwVan.GetWheel() != "BMW的轮胎" || bmwVan.GetEngine() != "BMW的引擎"{
		t.Error("benzSuv error")
	}

	//给我一辆混合车型
	complexCar := director.CreateComplexCar()
	if complexCar.GetWheel() != "benz的轮胎" || complexCar.GetEngine() != "BMW的引擎"{
		t.Error("benzSuv error")
	}
}
/**
	注意最后一个运行结果片段，我们可以立刻生产出一辆混合车型，只要有设计蓝图，这非常容易实现。反观我们的抽象工厂模式，它是不可能实现该功能的，
	因为它更关注的是整体，而不关注到底用的是奔驰引擎还是宝马引擎，而我们的建造者模式却可以很容易地实现该设计，市场信息变更了，我们就可以立刻跟进，生产出客户需要的产品。
 */