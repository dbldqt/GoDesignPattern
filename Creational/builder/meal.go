package builder

import (
	"fmt"
	"strconv"
)
//定义meal类
type Meal struct {
	items []IItem
}

func (meal *Meal) AddItem(item IItem){
	meal.items = append(meal.items,item)
}

func (meal *Meal) GetCost() float64{
	cost := 0.0
	for _,item := range meal.items {
		cost += item.Price()
	}
	return cost
}

func (meal *Meal) ShowItems(){
	for _,item := range meal.items{
		fmt.Println("Item : "+item.Name())
		fmt.Println("Packing : "+item.Packing().Pack())
		fmt.Println("Price : "+strconv.FormatFloat(item.Price(),'f',6,64))
	}
}






