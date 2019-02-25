package Interpreter

import (
	"strings"
)

//定义expression接口
type IExpression interface {
	Interpret(context string) bool
}
//定义Expression实体类
type TerminalExpression struct {
	data string
}

func (self *TerminalExpression) Interpret(context string) bool {
	return strings.Contains(context,self.data)
}

func NewTerminalExpression(data string)IExpression{
	return &TerminalExpression{data}
}
//定义OrExpression
type OrExpression struct {
	exp1 IExpression
	exp2 IExpression
}

func (self *OrExpression) Interpret(context string) bool {
	return self.exp1.Interpret(context) || self.exp2.Interpret(context)
}

func NewOrExpression(exp1,exp2 IExpression) IExpression{
	return &OrExpression{exp1,exp2}
}

//定义AndExpression
type AndExpression struct {
	exp1 IExpression
	exp2 IExpression
}

func (self *AndExpression) Interpret(context string) bool {
	return self.exp1.Interpret(context) && self.exp2.Interpret(context)
}

func NewAndExpression(exp1,exp2 IExpression)IExpression{
	return &AndExpression{exp1,exp2}
}













