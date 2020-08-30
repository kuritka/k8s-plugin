package runner

import (
	"github.com/kuritka/plugin/common/guard"
	"github.com/kuritka/plugin/common/log"
)

//CmdRunner is running all commands
type CmdRunner struct {
	service ICmdRunner
}

var logger = log.Log

//New creates new instance of CmdRunner
func New(command ICmdRunner) *CmdRunner {
	return &CmdRunner{
		command,
	}
}

//MustRun runs service once and panics if service is broken
func (r *CmdRunner) MustRun() {
	err := r.service.Run()
	guard.FailOnError(err, "command %s failed", r.service)
}
