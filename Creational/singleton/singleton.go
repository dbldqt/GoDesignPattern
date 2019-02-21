package singleton

import "sync"
var once sync.Once
type Singleton struct {

}
var singleton *Singleton

func GetSingleTon() *Singleton{
	once.Do(func(){
		singleton = &Singleton{}
	})
	return singleton
}

