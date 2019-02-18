package main
//定义接口
type IMethod interface {
	ValueMethod(name1 string)
	PointerMehtod(name2 string)
}
//父类
type StructParent struct {
	Name string
}

func (sp StructParent) M1(name1 string) {
	sp.Name = name1
}

func (sp *StructParent) M2(name2 string) {
	sp.Name = name2
}


type StructChild struct {
	StructParent
}

type PointerStructChild struct {
	*StructParent
}