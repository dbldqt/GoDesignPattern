package adapter

import (
	"fmt"
	"strings"
)

type IMediaPlayer interface {
	Play(audioType string,file string)
}

type AudioPlayer struct {

}

func (audio *AudioPlayer) Play(audioType string, file string) {
	lowerType := strings.ToLower(audioType)
	if lowerType == "mp3"{
		fmt.Println("audio player is playing "+file)
	}else if lowerType == "vlc" || lowerType == "mp4"{
		adapter := NewMediaAdapter(audioType)
		adapter.Play(audioType,file)
	}else{
		fmt.Println("not supported media type")
	}
}



