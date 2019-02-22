package strategy

type Context struct {
	Strategy IStrategy
}

func (con *Context) ExecStrategy(num1,num2 int) int{
	return con.Strategy.DoOperation(num1,num2)
}

func NewContext(strategy IStrategy) *Context{
	return &Context{strategy}
}
