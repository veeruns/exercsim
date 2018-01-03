//Package clock implements clock type and associated methods
package clock

import "fmt"

//Clock exports the clock structure
type Clock struct {
	hour   int
	minute int
}

//Add implements Add method for clock
func (cl Clock) Add(min int) Clock {

	cl.minute = cl.minute + min

	t1, t2 := rollover(cl.hour, cl.minute)
	cl.hour = t1
	cl.minute = t2
	return Clock{t1, t2}
}

//New constructor for Clock
func New(hour int, minute int) Clock {
	t1, t2 := rollover(hour, minute)
	return Clock{t1, t2}
}

//String method to print the time
func (cl Clock) String() string {
	var hour, minute int
	hour = cl.hour
	minute = cl.minute
	t1, t2 := rollover(hour, minute)
	var retstring string
	retstring = fmt.Sprintf("%.2d:%.2d", t1, t2)
	return retstring

}

//rollover method that is private to take care of rollingover
func rollover(hour int, minute int) (int, int) {
	var totalmin int
	var residualminute int
	totalmin = (hour*60 + minute) % 1440
	if totalmin < 0 {
		totalmin = totalmin + 1440

	}
	//fmt.Printf("Total minute %d\n", totalmin)
	ophour := totalmin / 60
	residualminute = totalmin % 60

	if ophour > 23 {
		ophour = ophour % 24
	}
	return ophour, residualminute

}
