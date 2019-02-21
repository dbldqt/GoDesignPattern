package adapter

import "testing"

func TestNewMediaAdapter(t *testing.T) {
	audioPlayer := AudioPlayer{}
	audioPlayer.Play("mp4","file_mp4")
	audioPlayer.Play("vlc","file_vlc")
	audioPlayer.Play("mp3","file_mp3")
	audioPlayer.Play("mp5","file_mp5")
}
