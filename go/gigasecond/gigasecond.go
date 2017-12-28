//Package gigasecond add 1e9 seconds to any date given to it
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond adds one 1e9 seconds to the input and returns.
// this returns the date after adding 1 giga seconds

func AddGigasecond(t time.Time) time.Time {
	var dura time.Duration = 1000000000 * time.Second
	return t.Add(dura)
}
