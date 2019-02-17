package builder
//定义打包接口
type Packing interface {
	Pack() string
}
//定义纸盒类
type Wrapper struct {

}

func (wrapper Wrapper) Pack() string {
	return "Wrapper"
}

//定义瓶子包装类
type Bottle struct {

}

func (bottle Bottle) Pack() string {
	return "Bottle"
}



