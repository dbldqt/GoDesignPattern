package decorator

type IShape interface {
	Draw()
}

type Rectangle struct {

}

func (rect * Rectangle) Draw(){
	println("Shape:rectangle")
}

type Circle struct {

}

func (cir *Circle) Draw(){

}

