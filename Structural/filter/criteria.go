package filter

import (
	"strings"
)

//定义标准
type ICriteria interface {
	meetCriteria(persons []Person) []Person
}
//定义男性过滤器
type CriteriaMale struct {

}

func (cm *CriteriaMale) meetCriteria(persons []Person) []Person {
	malePersons := make([]Person,len(persons))
	for _,person := range persons{
		if strings.ToLower(person.gender) == "male"{
			malePersons = append(malePersons,person)
		}
	}
	return malePersons
}

func NewCriteriaMale() *CriteriaMale{
	return &CriteriaMale{}
}
//定义女性过滤器
type CriteriaFemale struct {

}

func (cfm *CriteriaFemale) meetCriteria(persons []Person) []Person {
	femalePersons := make([]Person,len(persons))
	for _,person := range persons{
		if strings.ToLower(person.gender) == "female"{
			femalePersons = append(femalePersons,person)
		}
	}
	return femalePersons
}

func NewCriteriaFemale() *CriteriaFemale{
	return &CriteriaFemale{}
}
//定义单身过滤器
type CriteriaSingle struct {

}

func (sc *CriteriaSingle) meetCriteria(persons []Person) []Person {
	singlePersons := make([]Person,len(persons))
	for _,person := range persons{
		if strings.ToLower(person.maritalStatus) == "single"{
			singlePersons = append(singlePersons,person)
		}
	}
	return singlePersons
}

func NewCriteriaSingle() *CriteriaSingle{
	return &CriteriaSingle{}
}
//定义任意组合的标准类
type AndCriteria struct {
	criteria ICriteria
	otherCriteria ICriteria
}

func (ac *AndCriteria) meetCriteria(persons []Person) []Person {
	firstPersons := ac.criteria.meetCriteria(persons)
	secondePersons := ac.otherCriteria.meetCriteria(firstPersons)
	return secondePersons
}

func NewAndCriteria(criteria ICriteria,otherCriteria ICriteria) *AndCriteria{
	return &AndCriteria{
		criteria,
		otherCriteria,
	}
}



