package adapter

import (
	"fmt"
	"strings"
)

type MediaAdapter struct {
	advancedMediaPlayer IAdvancedMediaPlayer
}

func (mead *MediaAdapter) Play(fileType string, file string) {
	if strings.ToLower(fileType) == "vlc" {
		mead.advancedMediaPlayer.PlayVlc(file)
	}else if strings.ToLower(fileType) == "mp4"{
		mead.advancedMediaPlayer.PlayMp4(file)
	}else{
		fmt.Println("adapter not adapter")
	}
}

func NewMediaAdapter(fileType string) *MediaAdapter{
	if strings.ToLower(fileType) == "vlc" {
		return &MediaAdapter{&VlcPlayer{}}
	}else if strings.ToLower(fileType) == "mp4"{
		return &MediaAdapter{&Mp4Player{}}
	}else{
		return nil
	}
}


