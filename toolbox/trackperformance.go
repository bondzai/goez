// trackperformance.go
package toolbox

import (
	"log"
	"time"
)

// TrackPerformance returns a function that when called, logs the time
// taken since TrackPerformance was called.
func TrackPerformance(taskName string) func() {
	timeStart := time.Now()
	return func() {
		timeEnd := time.Now()
		log.Printf("%s took %v to run.\n", taskName, timeEnd.Sub(timeStart))
	}
}
