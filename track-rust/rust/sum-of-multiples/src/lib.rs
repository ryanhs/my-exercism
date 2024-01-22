pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    // println!("# Sum the multiples of all of {:?} which are less than {}", factors, limit);

    let mut arr: Vec<u32> = Vec::with_capacity(factors.len() * (limit as usize));

    for f in factors {
        for i in 1..limit {
            let multiple = f * i;
            
            if multiple >= limit {
                break;
            }

            // println!("  > factor({}) > {} = {}", f, i, multiple);
            arr.push(multiple);
        }
    }

    // sort
    arr.sort();
    arr.dedup();

    // println!("\n  > arr u: {:?}", arr);

    return arr.iter().sum::<u32>();
}
