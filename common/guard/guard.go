//Package guard panics when error occurs
package guard

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

//var logger = log.Log

//FailOnError panics when error occurs.
func FailOnError(err error, message string, args ...interface{}) {
	if err != nil {
		fmt.Println(emoji.Sprintf(":error: ", message, err.Error()))
	}
}
