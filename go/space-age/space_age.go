package space

// Planet is a string type
// if the test case planet was defined as string then we could have gotten away with String
type Planet string

// constant earthyear in seconds

const earthyear = 31557600
const testVersion = 1

// conversion table to number of seconds in each Planet

var ctable = map[Planet]float64{

	"Earth":   earthyear,
	"Mercury": 0.2408467 * earthyear,
	"Venus":   0.61519726 * earthyear,
	"Mars":    1.8808158 * earthyear,
	"Jupiter": 11.862615 * earthyear,
	"Saturn":  29.447498 * earthyear,
	"Uranus":  84.016846 * earthyear,
	"Neptune": 164.79132 * earthyear,
}

// Age function will convert and return number of years on a given planet
func Age(seconds float64, planet Planet) float64 {
	return (seconds / ctable[planet])
}
