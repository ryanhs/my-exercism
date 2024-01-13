package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	fa, faErr := calc.FodderAmount(numberOfCows)
	if faErr != nil {
		return 0, faErr
	}

	ff, ffErr := calc.FatteningFactor()
	if ffErr != nil {
		return 0, ffErr
	}

	return fa * ff / float64(numberOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calc, numberOfCows)
}

type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

func (ice InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.numberOfCows, ice.customMessage)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{
			numberOfCows:  numberOfCows,
			customMessage: "there are no negative cows",
		}
	case numberOfCows == 0:
		return &InvalidCowsError{
			numberOfCows:  numberOfCows,
			customMessage: "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
