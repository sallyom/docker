// +build !windows

package machine

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/systemd"
)

/*
Register Machine with systemd.  There is a potential race condition here
where the container could have exited before the call gets made.  This
call requires the container.Pid.
*/
func Register(name string, id string, pid int, root_directory string) error {
	err := systemd.RegisterMachine(name, id, pid, root_directory)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to determine unit of process") {
			return nil
		}
		logrus.Errorf("Unable to RegisterMachine %s for %s: %s", name, id, err)
		return err
	}
	return nil
}

func Terminate(name string) {
	systemd.TerminateMachine(name)
}
