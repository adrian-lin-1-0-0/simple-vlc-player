package utils

import (
	"github.com/adrian-lin-1-0-0/simple-vlc-player/vlc"
)

func setVolume(player *vlc.Player, volume int) {
	if volume > 100 {
		player.SetVolume(100)
		return
	}
	if volume < 0 {
		player.SetVolume(0)
		return
	}
	player.SetVolume(volume)
}

func VolumeUp10(player *vlc.Player) {
	setVolume(player, player.GetVolume()+10)
}

func VolumeDown10(player *vlc.Player) {
	setVolume(player, player.GetVolume()-10)
}

func setMediaTime(player *vlc.Player, milliseconds int) {

	length := player.GetMediaLength()
	if milliseconds > length {
		player.SetMediaTime(length)
		return
	}

	if milliseconds < 0 {
		player.SetMediaTime(0)
		return
	}

	player.SetMediaTime(milliseconds)
}

func Forward10Second(player *vlc.Player) {
	setMediaTime(player, player.GetMediaTime()+10*1000)
}

func Backward10Second(player *vlc.Player) {
	setMediaTime(player, player.GetMediaTime()-10*1000)
}
