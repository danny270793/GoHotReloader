package watcher

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Watcher struct {
	path string
}

func New(path string) Watcher {
	return Watcher{path: path}
}

func (w *Watcher) Read(c chan string) error {
	inotifyFd, err := unix.InotifyInit()
	if err != nil {
		return fmt.Errorf("inotify not created because: %v", err)
	}
	defer unix.Close(inotifyFd)

	watchDescriptor, err := unix.InotifyAddWatch(inotifyFd, w.path, unix.IN_MODIFY|unix.IN_CREATE|unix.IN_DELETE)
	if err != nil {
		return fmt.Errorf("path not added to inotify because: %v", err)
	}
	defer unix.InotifyRmWatch(inotifyFd, uint32(watchDescriptor))

	buffer := make([]byte, unix.SizeofInotifyEvent*10)
	for {
		n, err := unix.Read(inotifyFd, buffer)
		if err != nil {
			return fmt.Errorf("inotify not read because: %v", err)
		}

		var offset uint32
		for offset < uint32(n) {
			event := (*unix.InotifyEvent)(unsafe.Pointer(&buffer[offset]))
			nameBytes := buffer[offset+unix.SizeofInotifyEvent : offset+unix.SizeofInotifyEvent+event.Len]
			fileName := string(nameBytes[:len(nameBytes)-1])

			c <- fileName

			offset += unix.SizeofInotifyEvent + event.Len
		}
	}

	return nil
}
