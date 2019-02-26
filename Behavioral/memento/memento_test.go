package memento

import (
	"fmt"
	"testing"
)

func TestNewCareTaker(t *testing.T) {
	originator := Originator{}
	careTaker := NewCareTaker()

	originator.SetState("state #1")
	originator.SetState("state #2")

	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("state #3")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("state #4")

	fmt.Println("current State is "+originator.Getstate())
	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("first saved State is "+originator.Getstate())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("second saved State is "+originator.Getstate())
}
