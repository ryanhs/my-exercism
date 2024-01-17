pub fn square(s: u32) -> u64 {
    if s < 1 || s > 64 {
        panic!("Square must be between 1 and 64")
    }

    if s == 1 {
        return 1;
    }

    square(s - 1) * 2
}

pub fn total() -> u64 {
    let c64 = square(8 * 8);

    let checked = c64.checked_mul(2);

    if checked.is_none() {
        return u64::MAX;
    }

    checked.unwrap()
}
