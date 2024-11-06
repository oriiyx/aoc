use aoc::{convert_line_to_integer, process_file_line_by_line};

pub fn run() -> i32 {
    let lines_vec = process_file_line_by_line("./src/year_2019/input/01.txt", convert_line_to_integer);
    return lines_vec.unwrap()
        .iter()
        .map(|&x| recursive_fuel(x))
        .sum();
}

fn recursive_fuel(fuel: i32) -> i32 {
    let result = (fuel / 3) - 2;

    if result >= 0 {
        return result + recursive_fuel(result);
    }

    return 0;
}