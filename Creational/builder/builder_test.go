package builder

import "testing"

func TestMealBuilder_PrepareNonBVegMeal(t *testing.T) {
	mealBuilder := MealBuilder{}
	mealBuilder.PrepareNonBVegMeal()
	mealBuilder.PrepareVegMeal()
}