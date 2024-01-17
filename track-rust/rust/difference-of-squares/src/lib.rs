

// The square of the sum of the first ten natural numbers is
// (1 + 2 + ... + 10)² = 55² = 3025.
pub fn square_of_sum(n: u32) -> u32 {
    let mut a: u32 = 0;
    for x in 1..(n+1) {
        a += x;
    }

    a.pow(2)
}

// The sum of the squares of the first ten natural numbers is
// 1² + 2² + ... + 10² = 385.
pub fn sum_of_squares(n: u32) -> u32 {
    let mut a: u32 = 0;
    
    for x in 1..(n+1) {
        a += x.pow(2);
    }

    a
}

pub fn difference(n: u32) -> u32 {
    square_of_sum(n) - sum_of_squares(n)
}
