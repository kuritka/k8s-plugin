//Package guard panics when error occurs
package guard

import (
	"github.com/kuritka/plugin/common/log"
)

var logger = log.Log

//FailOnError panics when error occurs.
func FailOnError(err error, message string, v ...interface{}) {
	if err != nil {
		logger.Panic().Err(err).Msgf(message, v)
	}
}
