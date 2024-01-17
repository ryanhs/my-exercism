use std::u32;

pub fn is_armstrong_number(num: u32) -> bool {
    if num == 0 {
        return true;
    }

    println!("num: {}", num);
    println!("num is ok: {}", num < u32::MAX);

    let num_str = num.to_string();
    let try_armstrong = num_str
        .chars()
        .map(|d: char| d.to_digit(10).unwrap().pow(num_str.len() as u32))
        .try_fold(0, |a: u32, c| a.checked_add(c));
    
    // if overflow u32
    if try_armstrong.is_none() {
        return false;
    }

    try_armstrong.unwrap() == num
}

// 31_381_118_658
// 4_294_967_295
// 3_486_843_450
// 3_486_784_401