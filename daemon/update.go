package daemon

import (
	"github.com/docker/docker/engine"
)

func (daemon *Daemon) ContainerUpdate(job *engine.Job) engine.Status {
	if len(job.Args) != 1 {
		return job.Errorf("Usage: %s CONTAINER", job.Name)
	}
	name := job.Args[0]
	container, err := daemon.Get(name)
	if err != nil {
		return job.Error(err)
	}
	
	if err := daemon.execDriver.Update(container.command, 132321); err != nil {
		return job.Errorf("Cannot update container %s: %s", name, err)
	}
	container.LogEvent("update")
	return engine.StatusOK
}
