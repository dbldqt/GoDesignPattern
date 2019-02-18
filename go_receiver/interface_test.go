package go_receiver

import "testing"

func TestInterfaceReceiver(t *testing.T){
	var im IMethod  					 		 //接口变量
	sp := StructParent{"name0"}           //实现接口的结构体
	psp := &sp                                   //实现接口的结构体的指针
	sc := StructChild{sp}            //内嵌实现接口的匿名结构体的结构体
	psc := PointerStructChild{psp}   //内嵌实现接口的匿名结构体指针的结构体

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





	im = &sp
	im.PointerMehtod("name1")
	sc.PointerMehtod("name1")
	psc.PointerMehtod("name1")
}
func main(){
	//t2.M1() 相当于M1(t2)， t2 是指针类型，编译器自动转换，取 t2 的值并拷贝一份传给 M1。
	//t2.M2() 相当于M2(t2)，都是指针类型，不需要转换。
	//*T 类型的变量也是拥有这两个方法的。
	// */
	//m2 := &T{"name0"}
	//m2.M1("name1")
	//if m2.Name != "name0"{
	//}
	//m2.M2("name2")
	//if m2.Name != "name2"{
	//}
	////给定指针必定能得到变量，给定变量不一定能得到指针
	//
	//var interf Interf
	//m := T{"name0"}
	//
	////interf = m  m无法作为Interf的子类使用，因为T的M2是用的指针接受者实现的
	//interf = &m  //m的引用可以作为Interf的子类使用，虽然M1使用值接收者实现
	////给定指针必定能够找到变量，编译器会自动进项转换
	////给定变量不一定能获取变量的指针
	//interf.M1("name1")
	//
	//m := T{"name0"}
	//s := S{m}           //此处嵌入的并不是m本身，而是m的拷贝
	//sm := SM{&m}        //此处将m的指针嵌入，虽然也是指针的拷贝，但是该拷贝依然指向m
	//
	////将 T 嵌入 S， 那么 T 拥有的方法和属性 S 也是拥有的，但是接收者却不是 S 而是 T。
	////s.M1() 相当于 M1(t1) 而不是 M1(s)
	////最后 t1 的值没有改变，因为我们嵌入的是 T 类型，所以 S{t1} 的时候是将 t1 拷贝了一份。
	//
	//s.M1("name1")
	//sm.M1("name1")
	//
	//if s.Name != s.T.Name{
	//}
	//if sm.Name != sm.T.Name{
	//}
	//
	//if s.Name == "name1" || m.Name == "name1"{
	//}
	//if sm.Name == "name1" || m.Name == "name1"{
	//}
	//
	//
	//s.M2("name2")
	//sm.M2("name2")
	//
	//if s.Name != s.T.Name{
	//}
	//if sm.Name != sm.T.Name{
	//}
	//
	//if s.Name != "name2" || m.Name != "name2"{
	//}
	//if sm.Name != "name2" || m.Name != "name2"{
	//}
	//
	//m := T{"name0"}
	//s := S{m}
	////interf = s   s无法作为Interf的子类使用，因为T的M2是用的指针接受者实现的
	//interf = &s  //s的引用可以作为Interf的子类使用，这一点跟不嵌套的T一样
	//interf.M1("name1")
}
