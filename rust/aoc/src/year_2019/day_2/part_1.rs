use aoc::{convert_line_to_string, process_file_line_by_line};

pub fn run() {
    let lines_vec = process_file_line_by_line("./src/year_2019/input/02.txt", convert_line_to_string);
    let unwrap = lines_vec.unwrap();
    for line in unwrap.iter() {
        // Split each line by commas and parse into integers
        let mut split_values: Vec<i32> = line.split(',')
            .map(|s| s.trim().parse::<i32>().unwrap_or(0))
            .collect();
        let split_values_slice: &mut [i32] = &mut split_values[..];

        // Set the 1202 program alarm state
        split_values_slice[1] = 12;
        split_values_slice[2] = 2;

        // Process the intcode program
        let mut i = 0;
        while i < split_values_slice.len() {
            match split_values_slice[i] {
                1 => {
                    // Perform addition
                    let pos_1 = split_values_slice[i + 1] as usize;
                    let pos_2 = split_values_slice[i + 2] as usize;
                    let pos_i = split_values_slice[i + 3] as usize;
                    split_values_slice[pos_i] = split_values_slice[pos_1] + split_values_slice[pos_2];
                    i += 4;
                }
                2 => {
                    // Perform multiplication
                    let pos_1 = split_values_slice[i + 1] as usize;
                    let pos_2 = split_values_slice[i + 2] as usize;
                    let pos_i = split_values_slice[i + 3] as usize;
                    split_values_slice[pos_i] = split_values_slice[pos_1] * split_values_slice[pos_2];
                    i += 4;
                }
                99 => break,  // Program halt
                _ => {
                    i += 1;  // Move to the next instruction if opcode isn't 1, 2, or 99
                }
            }
        }

        // Print the result at position 0 after the program halts
        println!("Value at position 0: {}", split_values_slice[0]);
    }
}