
pub fn verse(n: u32) -> String {
    if n == 0 {
        return format!("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n");
    }

    if n == 1 {
        return format!("1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n");
    }

    if n == 2 {
        return format!("2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n");
    }

    return format!("{} bottles of beer on the wall, {} bottles of beer.\nTake one down and pass it around, {} bottles of beer on the wall.\n", n, n, n - 1);
}

pub fn sing(start: u32, end: u32) -> String {
    let mut acc = String::new();

    for i in (end..(start + 1)).rev() {
        if i != start {
            acc.push_str("\n");
        }
        acc.push_str(&verse(i));
    }

    return acc
}
