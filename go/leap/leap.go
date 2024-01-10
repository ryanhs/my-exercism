// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// https://learn.microsoft.com/en-us/office/troubleshoot/excel/determine-a-leap-year
// 1. If the year is evenly divisible by 4, go to step 2. Otherwise, go to step 5.
// 2. If the year is evenly divisible by 100, go to step 3. Otherwise, go to step 4.
// 3. If the year is evenly divisible by 400, go to step 4. Otherwise, go to step 5.
// 4. The year is a leap year (it has 366 days). > return true
// 5. The year is not a leap year (it has 365 days). > return false

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	step4 := true
	step5 := false

	// 1
	if year % 4 != 0 {
		return step5
	}

	// 2
	if year % 100 != 0 {
		return step4
	}

	// 3
	if year % 400 == 0 {
		return step4
	}

	return step5;
}
