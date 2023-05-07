package utils

import "github.com/adrian-lin-1-0-0/simple-vlc-player/vlc"

type MediaPlayer struct {
	Instance *vlc.Instance
	Media    *vlc.Media
	Player   *vlc.Player
}

func NewMediaPlayer(fileName *string) *MediaPlayer {
	vlcInstance, err := vlc.Init("--no-video", "--quiet")
	if err != nil {
		panic(err)
	}

	media, err := vlcInstance.NewMedia(*fileName)
	if err != nil {
		panic(err)
	}

	player, err := vlcInstance.NewPlayer()
	if err != nil {
		panic(err)
	}

	err = player.SetMedia(media)
	if err != nil {
		panic(err)
	}

	return &MediaPlayer{
		Instance: vlcInstance,
		Media:    media,
		Player:   player,
	}
}

func (m *MediaPlayer) Release() {
	m.Media.Release()
	m.Player.Release()
	m.Instance.Release()
}

func (m *MediaPlayer) SetVolume(volume int) {
	if volume > 100 {
		m.Player.SetVolume(100)
		return
	}
	if volume < 0 {
		m.Player.SetVolume(0)
		return
	}
	m.Player.SetVolume(volume)
}

func (m *MediaPlayer) VolumeUp10() {
	m.SetVolume(m.Player.GetVolume() + 10)
}

func (m *MediaPlayer) VolumeDown10() {
	m.SetVolume(m.Player.GetVolume() - 10)
}

func (m *MediaPlayer) SetMediaTime(milliseconds int) {

	length := m.Player.GetMediaLength()
	if milliseconds > length {
		m.Player.SetMediaTime(length)
		return
	}

	if milliseconds < 0 {
		m.Player.SetMediaTime(0)
		return
	}

	m.Player.SetMediaTime(milliseconds)
}

func (m *MediaPlayer) Forward10Second() {
	m.SetMediaTime(m.Player.GetMediaTime() + 10*1000)
}

func (m *MediaPlayer) Backward10Second() {
	m.SetMediaTime(m.Player.GetMediaTime() - 10*1000)
}

func (m *MediaPlayer) Play() error {
	return m.Player.Play()
}
