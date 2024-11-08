use std::fs::File;
use std::io::{self, BufRead};

pub fn process_file_line_by_line<F, T>(filename: &str, mut process_line: F) -> io::Result<Vec<T>>
where
    F: FnMut(&str) -> T,
{
    let file = File::open(filename)?;
    let reader = io::BufReader::new(file);
    let mut results = Vec::new();

    for line in reader.lines() {
        let line = line?; // Each line as a Result<String>
        results.push(process_line(&line)); // Apply the provided function to each line
    }

    Ok(results)
}

pub fn convert_line_to_integer(line: &str) -> i32 {
    line.trim().parse::<i32>().unwrap_or(0)
}

pub fn convert_line_to_string(line: &str) -> String {
    line.trim().to_string()
}