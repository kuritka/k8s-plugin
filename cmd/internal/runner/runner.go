package runner

import (
	"github.com/kuritka/plugin/common/guard"
	"github.com/kuritka/plugin/common/log"
)

type cmdRunner struct {
	service ICmdRunner
}

var logger = log.Log

func New(command ICmdRunner) *cmdRunner {
	return &cmdRunner{
		command,
	}
}

//Run service once and panics if service is broken
func (r *cmdRunner) MustRun() {
	logger.Info().Msgf("command %s started", r.service)
	err := r.service.Run()
	guard.FailOnError(err, "command %s failed", r.service)
}
