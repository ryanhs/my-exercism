package collatzconjecture

// import "fmt"
import "errors"

// Take any positive integer n.
// If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1.
// Repeat the process indefinitely.
// The conjecture states that no matter which number you start with, you will always reach 1 eventually.
// Given a number n, return the number of steps required to reach 1.


func Processor(n int, i int) (int, int) {
	// final
	if n <= 1 {
		return n, i;
	}

	var newI = i + 1;
	// fmt.Printf("%d. %d\n", newI, n);
	
	// If n is even, divide n by 2 to get n / 2.
	if n % 2 == 0 {
		return Processor(n / 2, newI);
	}
	
	// If n is odd, multiply n by 3 and add 1 to get 3n + 1.
	return Processor((3 * n) + 1, newI);
}


func CollatzConjecture(n int) (int, error) {
	// if 1 = 0 step, but not error
	if n == 1 {
		return 0, nil;
	}

	var _, res = Processor(n, 0);

	// 0 step? errors
	if res == 0 {
		return n, errors.New("0 steps")
	}

	return res, nil;
}
