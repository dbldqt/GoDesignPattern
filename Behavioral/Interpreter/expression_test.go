package Interpreter

import (
	"fmt"
	"strconv"
	"testing"
)

//规则：Robert 和 John 是男性
func GetMaleExpression() IExpression{
	robert := TerminalExpression{"Robert"}
	john := TerminalExpression{"John"}
	return NewOrExpression(&robert,&john)
}
//规则：Julie 是一个已婚的女性
func GetMarriedWomanExpression() IExpression{
	julie := TerminalExpression{"Julie"}
	married := TerminalExpression{"Married"}
	return NewAndExpression(&julie,&married)
}

func TestTerminalExpression_Interpret(t *testing.T) {
	isMale := GetMaleExpression()
	isMarriedWWoman := GetMarriedWomanExpression()

	fmt.Println("John is male? "+strconv.FormatBool(isMale.Interpret("John")))
	fmt.Println("Julie is a married women? "+strconv.FormatBool(isMarriedWWoman.Interpret("Married Julie")))
}

