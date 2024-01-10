// Package weather provides tools to forecast.
// forcast a b c.
package weather

// CurrentCondition dwa nd wa.
var CurrentCondition string

// CurrentLocation dawdaw.
var CurrentLocation string

// Forecast dwa dwafwad.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
