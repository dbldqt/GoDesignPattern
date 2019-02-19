package go_receiver

import "testing"

func TestInterfaceReceiver(t *testing.T){
	var im IMethod  					 		 //接口变量
	sp := StructParent{"name0"}           //实现接口的结构体
	psp := &sp                                   //实现接口的结构体的指针
	sc := StructChild{sp}            //内嵌实现接口的匿名结构体的结构体
	psc := PointerStructChild{psp}   //内嵌实现接口的匿名结构体指针的结构体

	//go语言默认所有的参数都是值拷贝传递,
	//当参数为值类型则会赋值一份，参数为指针虽然也会复制一份，但是复制的指针已然只想原来的变量
	//接收者可以看作是函数的第一个参数，
	//即这样的： func M1(t T), func M2(t *T)。
	//编译器在需要的时候会自动进行sp和psp之间的转换，但是请注意，psp任何情况下均可以转换为sp，但是sp并不总能获得其地址并转换为psp

	//当调用 sp.ValueMethod() 时相当于 ValueMethod(sp) ，实参和行参都是类型StructParent，可以接受。
	//此时在ValueMethod()中的sp只是sp的值拷贝，所以ValueMethod()的修改影响不到sp
	sp.ValueMethod("name1")
	if sp.Name == "name1"{
		t.Error("sp.ValueMethod error "+sp.Name)
	}

	//使用结构体的指针一样可以调用值接受者实现的方法，此时编译器会转换为*psp,一样无法改变对象的Name值
	psp.ValueMethod("name1")
	if psp.Name == "name1"{
		t.Error("psp.ValueMethod error "+psp.Name)
	}

	//当调用 sp.PointerMehtod() 相当于PointerMehtod(&sp)，这时编译器自动将sp类型传给了&sp，
	//此时编译器依然会复制&sp的值作为参数传递给PointerMethod(),只是复制的是指针的值，指针依然只想sp，所以sp的值会被改变
	sp.PointerMehtod("name2")
	if sp.Name != "name2"{
		t.Error("sp.PointerMethod error "+sp.Name)
	}

	//使用psp调用时，因为跟生命的接受者一致所以编译器无需进行转换，其结果跟使用sp调用一致
	psp.PointerMehtod("name2")
	if psp.Name != "name2"{
		t.Error("psp.PointerMethod error "+psp.Name)
	}

	//重置sp的值为初始值，便于后续测试
	sp.PointerMehtod("name0")

	//此处因为psp是sp的指针，所以这两个的状态始终同步
	if sp.Name != psp.Name{
		t.Error("sp not equal psp error "+sp.Name+" "+psp.Name)
	}
	sp.PointerMehtod("name2")
	if sp.Name != psp.Name{
		t.Error("sp not equal psp error "+sp.Name+" "+psp.Name)
	}
	psp.PointerMehtod("name1")
	if sp.Name != psp.Name{
		t.Error("sp not equal psp error "+sp.Name+" "+psp.Name)
	}

	//重置sp的值为初始值，便于后续测试
	sp.PointerMehtod("name0")

	sp.PointerMehtod("name1")
	//此处表明，虽然sp作为sc的匿名内置对象，但是sc中的sp并不是最初的sp，而是sp的一个副本，
	if sp.Name == sc.Name{
		t.Error("sp.Name equal sc.Name error"+sp.Name+" "+sc.Name)
	}
	//此处表明psc中的psp虽然是psp的复本，但是跟psp一样都只想sp，所以值同步
	if sp.Name != psc.Name{
		t.Error("sp.Name not equal psc.Name error"+sp.Name+" "+sc.Name)
	}

	//重置sp的值为初始值，便于后续测试
	sp.PointerMehtod("name0")
	//下面的测试再次证明上面的情况
	sc.PointerMehtod("name1")
	if sp.Name == sc.Name{
		t.Error("sp.Name equal sc.Name error"+sp.Name+" "+sc.Name)
	}
	psc.PointerMehtod("name2")
	if sp.Name != psc.Name{
		t.Error("sp.Name not equal psc.Name error"+sp.Name+" "+sc.Name)
	}

	//im = sp sp不能赋值给接口变量，因为sp并未完全实现该接口（变量并不总能获得其指针）
	im = psp  //psp可以赋值给接口变量，因为通过变量指针总能获得其变量
	//im = sc   sc同sp一样不能赋值给im
	im = psc  //psc同psp一样可以赋值给im
	im.ValueMethod("name0")
}
