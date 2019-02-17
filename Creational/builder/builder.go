package builder

//定义MealBuilder类，该类用于构建Meal对象
type MealBuilder struct {

}

func (mb *MealBuilder) PrepareVegMeal(){
	meal := Meal{
	}
	meal.items = make([]IItem,5)
	meal.AddItem(VegBurger{})
	meal.AddItem(Coke{})
}

func (mb *MealBuilder) PrepareNonBVegMeal(){
	meal := Meal{
	}
	meal.items = make([]IItem,5)
	meal.AddItem(ChickenBurger{})
	meal.AddItem(Pepsi{})
}
