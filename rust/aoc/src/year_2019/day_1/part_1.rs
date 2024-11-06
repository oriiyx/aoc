use aoc::{convert_line_to_integer, process_file_line_by_line};

pub fn run() -> i32 {
    let lines_vec = process_file_line_by_line("./src/year_2019/input/01.txt", convert_line_to_integer);
    return lines_vec.unwrap()
        .iter()
        .map(|&x| { (x / 3) - 2 })
        .sum();
}
