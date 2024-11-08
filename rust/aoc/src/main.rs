use std::env;
mod year_2019;


fn main() {
    let current_dir = env::current_dir().expect("Failed to get current directory");
    println!("Current working directory: {}", current_dir.display());

    let day1p1 = year_2019::day_1::part_1::run();
    println!("Day1 Part1: {:?}", day1p1);
    let day1p2 = year_2019::day_1::part_2::run();
    println!("Day1 Part2: {:?}", day1p2);

    println!("------");

    let day2p1 = year_2019::day_2::part_1::run();
    println!("Day2 Part1: {:?}", day2p1);
}
