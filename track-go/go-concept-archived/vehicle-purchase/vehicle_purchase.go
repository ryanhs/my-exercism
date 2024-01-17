package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1 string, option2 string) string {
	betterChoice := option1

	if strings.Compare(option1, option2) == 1 {
		betterChoice = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", betterChoice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new.
	// If it is at least 10 years old, it costs 50%.
	// If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.

	if age < 3 {
		return originalPrice * 0.8
	}

	if age >= 3 && age < 10 {
		return originalPrice * 0.7
	}

	return originalPrice * 0.5
}
