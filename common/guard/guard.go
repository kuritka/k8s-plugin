//Package guard panics when error occurs
package guard

import (
	"fmt"
	"github.com/kuritka/plugin/common/log"
)

var logger = log.Log

//FailOnError panics when error occurs.
func FailOnError(err error, message string, args ...interface{}) {
	if err != nil {
		logger.Fatal().Err(err).Msgf(message, args...)
	}
}

func HandleError(err error) {
	if err != nil {
		fmt.Printf("printer error: %s", err.Error())
	}
}
