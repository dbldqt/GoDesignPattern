package bridge

import "testing"

func TestNewCircle(t *testing.T) {
	redCircle := NewCircle(&RedCircle{})
	greenCircle := NewCircle(&GreenCircle{})

	redCircle.Draw()
	greenCircle.Draw()
}
