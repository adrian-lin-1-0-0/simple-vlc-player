package vlc

// #cgo LDFLAGS: -lvlc
// #cgo CFLAGS: -w
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import "errors"

type Player struct {
	player *C.libvlc_media_player_t
}

func (p *Player) Release() {
	if p.player == nil {
		return
	}
	C.libvlc_media_player_release(p.player)
}

func (p *Player) Stop() {
	C.libvlc_media_player_stop(p.player)
}

func (p *Player) Play() error {
	if C.libvlc_media_player_play(p.player) != 0 {
		return errors.New("libvlc_media_player_play failed")
	}
	return nil
}

func (p *Player) ReleaseMedia() {
	C.libvlc_media_player_release(p.player)
}

func (p *Player) SetMedia(m *Media) error {
	if m.media == nil {
		return errors.New("media is nil")
	}
	C.libvlc_media_player_set_media(p.player, m.media)
	return nil
}

func (p *Player) SetVolume(volume int) {
	C.libvlc_audio_set_volume(p.player, C.int(volume))
}

func (p *Player) GetVolume() int {
	return int(C.libvlc_audio_get_volume(p.player))
}

func (p *Player) GetMediaLength() int {
	return int(C.libvlc_media_player_get_length(p.player))
}

func (p *Player) GetMediaTime() int {
	return int(C.libvlc_media_player_get_time(p.player))
}

func (p *Player) SetMediaTime(time int) {
	C.libvlc_media_player_set_time(p.player, C.libvlc_time_t(time))
}
