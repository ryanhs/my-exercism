pub fn reverse(input: &str) -> String {
    println!("input str: {}", input);

    let chars: std::str::Chars<'_> = input.chars();
    
    let reversed: String = chars.rev().collect();

    reversed
}
