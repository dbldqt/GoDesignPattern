package observer

import "testing"

func TestBinaryObserver_Update(t *testing.T) {
	subject := NewSubject()
	NewObserver("binary",subject)
	NewObserver("octal",subject)
	NewObserver("hex",subject)
	subject.SetState(1)
}
