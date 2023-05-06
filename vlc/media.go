package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <stdlib.h>
// #include <string.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Media struct {
	media *C.libvlc_media_t
}

func (i *Instance) NewMedia(path string) (*Media, error) {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	media := C.libvlc_media_new_path(i.handle, cPath)

	if media == nil {
		return nil, errors.New("libvlc_media_new_path failed")
	}
	return &Media{media: media}, nil
}

func (m *Media) Release() {
	if m.media == nil {
		return
	}
	C.libvlc_media_release(m.media)
}
