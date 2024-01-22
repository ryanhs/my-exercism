fn is_prime(x: u32) -> bool {
    if x <= 1 {
        return false;
    }

    let sqrt_x = (x as f64).sqrt() as u32;
    for i in 2..=sqrt_x {
        if x % i == 0 {
            return false;
        }
    }

    true
}

pub fn nth(n: u32) -> u32 {
    if n == 0 {
        return 2;
    }

    let mut i: u32 = 1;
    let mut x: u32 = 3;

    loop {
        if is_prime(x) {
            if i == n {
                return x;
            }

            i += 1;
        }

        x += 2; // only odd numbers will be prime
    }
}
