package observer

import (
	"fmt"
	"strconv"
)

//定义观察者基类
type IObserver interface {
	Update()
}

type Observer struct {
	subject *Subject
}
func (ob *Observer) Update(){
	panic("this func need overwrited")
}

//定义实际的观察者类
type BinaryObserver struct {
	Observer
}

func (binOb *BinaryObserver) Update(){
	fmt.Println("BinaryObserver Update "+strconv.Itoa(binOb.subject.GetState()))
}

type OctalObserver  struct {
	Observer
}

func (binOb *OctalObserver ) Update(){
	fmt.Println("OctalObserver Update "+strconv.Itoa(binOb.subject.GetState()))
}

type HexaObserver struct {
	Observer
}

func (binOb *HexaObserver) Update(){
	fmt.Println("HexaObserver Update "+strconv.Itoa(binOb.subject.GetState()))
}

func NewObserver(obName string,subject *Subject) IObserver{
	switch obName{
	case "binary":
		observer := &BinaryObserver{Observer{subject}}
		observer.subject.Attach(observer)
		return observer
	case "octal":
		observer := &OctalObserver{Observer{subject}}
		observer.subject.Attach(observer)
		return observer
	case "hex":
		observer := &HexaObserver{Observer{subject}}
		observer.subject.Attach(observer)
		return observer
	default:
		return nil
	}
}


