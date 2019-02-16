package abstract_factory

import "testing"

func TestAbstrctFactory(t *testing.T) {
	sportDress := NewDress(SportDressFactory{})
	if sportDress.Cloth.GetClothStyle() != "sportCloth" || sportDress.Bag.GetBagStyle() != "sportBag"{
		t.Error("sport dress fail")
	}

	casualDress := NewDress(CasualDressFactory{})
	if casualDress.Cloth.GetClothStyle() != "casualCloth" || casualDress.Bag.GetBagStyle() != "casualBag" {
		t.Error("casual dress fail")
	}
}