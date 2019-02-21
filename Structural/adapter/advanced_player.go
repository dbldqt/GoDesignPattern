package adapter

import "fmt"

type IAdvancedMediaPlayer interface {
	PlayVlc(file string)
	PlayMp4(file string)
}


type VlcPlayer struct {

}

func (vlc *VlcPlayer) PlayVlc(file string) {
	fmt.Println("Vlc player is play "+file)
}

func (vlc *VlcPlayer) PlayMp4(file string) {

}


type Mp4Player struct {

}

func (mp4 *Mp4Player) PlayVlc(file string) {

}

func (mp4 *Mp4Player) PlayMp4(file string) {
	fmt.Println("Mp4 player is play "+file)
}
