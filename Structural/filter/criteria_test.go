package filter

import (
	"fmt"
	"testing"
)

func TestNewCriteriaMale(t *testing.T) {
	persons := make([]Person,10)
	persons = append(persons,NewPerson("Robert","Male", "Single"))
	persons = append(persons,NewPerson("John","Male", "Married"))
	persons = append(persons,NewPerson("Laura","Female", "Married"))
	persons = append(persons,NewPerson("Diana","Female", "Single"))
	persons = append(persons,NewPerson("Mike","Male", "Single"))
	persons = append(persons,NewPerson("Bobby","Male", "Single"))

	maleCri := NewCriteriaMale()
	//femaleCri := NewCriteriaFemale()
	//singleCri := NewCriteriaSingle()
	//andCrit := NewAndCriteria(singleCri,maleCri)

	fmt.Println("males:")
	printPersons(maleCri.meetCriteria(persons))

	//fmt.Println("females:")
	//printPersons(femaleCri.meetCriteria(persons))
	//
	//fmt.Println("single males:")
	//printPersons(andCrit.meetCriteria(persons))
}

func printPersons(persons []Person){
	for _,person := range persons{
		fmt.Println(person)
		fmt.Println(person.GetName()+" "+person.GetGender()+" "+person.GetMaritalStatus())
	}
}