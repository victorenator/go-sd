package sd

import "fmt"
import "net"
import "os"
import "strconv"
import "strings"
import "syscall"

const START_FD int = 3

type FileDescriptor struct {
	fd int
	name string
}

type FD interface {
        FD() int
        Name() string
	Listener() (net.Listener, error)
}

func (l FileDescriptor) FD() int {
    return l.fd
}

func (l FileDescriptor) Name() string {
    return l.name
}

func (l FileDescriptor) Listener() (net.Listener, error) {
	return net.FileListener(os.NewFile(uintptr(l.fd), l.name))
}

func ListenPID() int {
	pid, err := strconv.ParseInt(os.Getenv("LISTEN_PID"), 10, 16)
	if err != nil {
		return 0
	} else {
		return int(pid)
	}
}

func ListenFDs() int {
        if ListenPID() != os.Getpid() {
            return 0
        }
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

func Listeners() []FileDescriptor {
	fds := ListenFDs()
	names := ListenFDNames()
	ls := make([]FileDescriptor, fds)
	for i := 0; i < fds; i++ {
            ls[i].fd = START_FD + i
            ls[i].name = names[i]
	}

	return ls
}

func Notify(state string, fds ...int) {
	ns := os.Getenv("NOTIFY_SOCKET")
        fmt.Printf("Notify %v %v %s\n", state, fds, ns)
	if ns == "" {
		return
	}

	addr := &net.UnixAddr{
		Name: ns,
		Net:  "unixgram",
	}

	conn, err := net.DialUnix("unixgram", nil, addr)
	if err != nil {
		return
	}

	defer conn.Close()
        
        if len(fds) > 0 {
            rights := syscall.UnixRights(fds...)

            conn.WriteMsgUnix([]byte(state), rights, nil)
            
        } else {
            conn.Write([]byte(state))
        }
}
