package singleton

import "sync"

type Singleton struct {

}
var singleton *Singleton

func GetSingleTon() *Singleton{
	sync.Once{}.Do(func(){
		singleton = &Singleton{}
	})
	return singleton
}

