//Pacakge clock implements clock type and associated methods
package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func (cl *Clock) Add(min int) *Clock {
	var residueminute int
	fmt.Printf("Add Hour %d:minute %d Add %d\n", cl.hour, cl.minute, min)

	cl.minute = cl.minute + min
	fmt.Printf("Current minute %d\n", cl.minute)
	if cl.minute >= 60 {

		residueminute = cl.minute / 60
		cl.minute = cl.minute % 60
	}
	cl.hour = cl.hour + residueminute
	if cl.hour > 23 {
		cl.hour = cl.hour % 24
	}
	return cl
}

func New(hour int, minute int) *Clock {

	//var residueminute int
	cl := new(Clock)
	t1, t2 := rollover(hour, minute)
	fmt.Printf("Rollover NEW %.2d:%.2d\n", t1, t2)
	/*	fmt.Printf("Before Input Hour %d:minute %d\n", hour, minute)
		if hour < 0 {
			hour = 24 - (hour*-1)%24
		}
		if minute < 0 {
			minute = 60 - minute
		}
		fmt.Printf("After Hour %d:minute %d\n", hour, minute)

		cl.hour = hour
		cl.minute = minute
		if cl.minute >= 60 {

			residueminute = cl.minute / 60
			cl.minute = cl.minute % 60
		}
		fmt.Printf("Residue minute %d\n", residueminute)
		cl.hour = cl.hour + residueminute
		if cl.hour > 23 {
			cl.hour = cl.hour % 24
		}
		fmt.Printf("Hour %d:minute %d\n", cl.hour, cl.minute)
	*/
	cl.hour = t1
	cl.minute = t2
	return cl
}

func (cl *Clock) String() string {
	//var residueminute int
	//var convertminute int
	var hour, minute int
	hour = cl.hour
	minute = cl.minute
	t1, t2 := rollover(hour, minute)
	fmt.Printf("Rollover %.2d:%.2d\n", t1, t2)
	/*if hour < 0 {
		hour = 24 - (hour*-1)%24
	}
	if minute < 0 {
		minute = 60 - minute
	}
	cl.hour = hour
	cl.minute = minute
	if cl.minute >= 60 {

		residueminute = cl.minute / 60
		cl.minute = 0
	}
	cl.hour = cl.hour + residueminute
	if cl.hour > 23 {
		cl.hour = cl.hour % 24
	}
	*/

	var retstring string
	retstring = fmt.Sprintf("%.2d:%.2d", t1, t2)
	fmt.Printf("return string is %s\n", retstring)
	return retstring

}

func rollover(hour int, minute int) (int, int) {
	var totalmin int
	var residualminute int

	fmt.Printf("\nRollover Input Hour %d:minute %d\n", hour, minute)
	totalmin = (hour*60 + minute) % 1440
	if totalmin < 0 {
		fmt.Printf("Calculation %d, %d \n", totalmin, (totalmin*-1)%1440)
		totalmin = totalmin + 1440

	}
	fmt.Printf("Total minute %d\n", totalmin)
	ophour := totalmin / 60
	residualminute = totalmin % 60

	if ophour > 23 {
		ophour = ophour % 24
	}
	return ophour, residualminute

}
