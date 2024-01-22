pub fn build_proverb(list: &[&str]) -> String {
    // empty
    if list.len() == 0 {
        return String::new();
    }

    let mut output = String::new();

    for i in 0..list.len() {
        if i == (list.len() - 1) {
            output.push_str(
                format!("And all for the want of a {}.", list[0]).as_str(),
            );
            return output;
        }

        output.push_str(
            format!(
                "For want of a {} the {} was lost.\n",
                list[i],
                list[i + 1]
            )
            .as_str(),
        )
    }

    output
}
