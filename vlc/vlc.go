package vlc

// #cgo LDFLAGS: -lvlc
// #cgo CFLAGS: -w
// #include <vlc/vlc.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Instance struct {
	handle *C.libvlc_instance_t
}

func Init(args ...string) (*Instance, error) {

	argc := len(args)
	argv := make([]*C.char, argc)

	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	defer func() {
		for i := range argv {
			C.free(unsafe.Pointer(argv[i]))
		}
	}()

	handle := C.libvlc_new(C.int(argc), *(***C.char)(unsafe.Pointer(&argv)))
	if handle == nil {
		return nil, errors.New("libvlc_new failed")
	}

	return &Instance{
		handle: handle,
	}, nil
}

func (i *Instance) Release() {
	if i.handle == nil {
		return
	}
	C.libvlc_release(i.handle)
}

func (i *Instance) NewPlayer() (*Player, error) {
	player := C.libvlc_media_player_new(i.handle)
	if player == nil {
		return nil, errors.New("libvlc_media_player_new failed")
	}
	return &Player{player: player}, nil
}
