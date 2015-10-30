package sd

import "net"
import "os"
import "strconv"
import "strings"

const START_FD int = 3

type Listener struct {
	net.Listener
	name string
}

type Named interface {
	Name() string
}

func (l Listener) Name() string {
	return l.name
}

func ListenFDs() int {
	fds, err := strconv.ParseInt(os.Getenv("LISTEN_FDS"), 10, 16)
	if err != nil {
		return 0
	} else {
		return int(fds)
	}
}

func ListenFDNames() []string {
	return strings.Split(os.Getenv("LISTEN_FDNAMES"), ":")
}

func Listeners() ([]Listener, error) {
	fds := ListenFDs()
	names := ListenFDNames()
	ls := make([]Listener, fds)
	for i := 0; i < fds; i++ {
		fl, err := net.FileListener(os.NewFile(uintptr(START_FD+i), names[i]))
		if err == nil {
			ls[i].Listener = fl
			ls[i].name = names[i]

		} else {
			return nil, err
		}
	}

	return ls, nil
}
