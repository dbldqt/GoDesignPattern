package main

import "testing"

var interf Interf
var
func main(){
	/**
		接收者可以看作是函数的第一个参数，
		即这样的： func M1(t T), func M2(t *T)。
		go不是完全的面向对象的语言，所以用那种看起来像面向对象的语法来理解可能有偏差。

		当调用 t1.M1() 时相当于 M1(t1) ，实参和行参都是类型 T，可以接受。
		此时在M1()中的t只是t1的值拷贝，所以M1()的修改影响不到t1。

		当调用 t1.M2() => M2(t1)，这时编译器自动将 T 类型传给了 *T 类型，
		go可能会取 t1 的地址传进去： M2(&t1)。所以 M2() 的修改可以影响 t1 。
 */
	m1 := T{"name0"}
	m1.M1("name")
	if m1.Name == "name"{     //M1是值接受者，无法改变本对象的值

	}
	m1.M2("name2")
	if m1.Name != "name2"{    //M2是指针接受者，
	}

	/*
	t2.M1() 相当于M1(t2)， t2 是指针类型，编译器自动转换，取 t2 的值并拷贝一份传给 M1。
	t2.M2() 相当于M2(t2)，都是指针类型，不需要转换。
	*T 类型的变量也是拥有这两个方法的。
	 */
	m2 := &T{"name0"}
	m2.M1("name1")
	if m2.Name != "name0"{
	}
	m2.M2("name2")
	if m2.Name != "name2"{
	}
	//给定指针必定能得到变量，给定变量不一定能得到指针

	var interf Interf
	m := T{"name0"}

	//interf = m  m无法作为Interf的子类使用，因为T的M2是用的指针接受者实现的
	interf = &m  //m的引用可以作为Interf的子类使用，虽然M1使用值接收者实现
	//给定指针必定能够找到变量，编译器会自动进项转换
	//给定变量不一定能获取变量的指针
	interf.M1("name1")

	m := T{"name0"}
	s := S{m}           //此处嵌入的并不是m本身，而是m的拷贝
	sm := SM{&m}        //此处将m的指针嵌入，虽然也是指针的拷贝，但是该拷贝依然指向m

	//将 T 嵌入 S， 那么 T 拥有的方法和属性 S 也是拥有的，但是接收者却不是 S 而是 T。
	//s.M1() 相当于 M1(t1) 而不是 M1(s)
	//最后 t1 的值没有改变，因为我们嵌入的是 T 类型，所以 S{t1} 的时候是将 t1 拷贝了一份。

	s.M1("name1")
	sm.M1("name1")

	if s.Name != s.T.Name{
	}
	if sm.Name != sm.T.Name{
	}

	if s.Name == "name1" || m.Name == "name1"{
	}
	if sm.Name == "name1" || m.Name == "name1"{
	}


	s.M2("name2")
	sm.M2("name2")

	if s.Name != s.T.Name{
	}
	if sm.Name != sm.T.Name{
	}

	if s.Name != "name2" || m.Name != "name2"{
	}
	if sm.Name != "name2" || m.Name != "name2"{
	}

	m := T{"name0"}
	s := S{m}
	//interf = s   s无法作为Interf的子类使用，因为T的M2是用的指针接受者实现的
	interf = &s  //s的引用可以作为Interf的子类使用，这一点跟不嵌套的T一样
	interf.M1("name1")
}
