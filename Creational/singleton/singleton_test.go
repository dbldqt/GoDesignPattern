package singleton

import "testing"

func TestGetSingleTon(t *testing.T) {
	if GetSingleTon() != GetSingleTon(){
		t.Error("singleton error")
	}
}
