package strategy
//定义策略接口
type IStrategy interface {
	DoOperation (num1,num2 int) int
}

//定义加法策略类
type OperationAdd struct {

}
func (opadd *OperationAdd) DoOperation(num1, num2 int) int {
	return num1+num2
}

//定义减法策略类
type OperationSub struct {

}
func (opsub *OperationSub) DoOperation(num1, num2 int) int {
	return num1-num2
}

//定义乘法操作类
type OperationMultiply struct {

}

func (opmul *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1*num2
}






