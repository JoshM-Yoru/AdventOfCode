use std::{fs, num::ParseIntError};

fn main() -> Result<(), ParseIntError>{

    let input = fs::read_to_string("input.txt").unwrap();

    let mut elves = vec![];

    for elf in input.replace("\r\n", "\n").split("\n\n") {

        let mut total = 0;

        for calories in elf.lines() {
            let value = calories.parse::<u64>()?;

            total += value;
        }

        elves.push(total);

    }

    elves.sort_unstable();

    println!("Most calories: {}", elves.last().unwrap());

    println!("{}", elves.into_iter().rev().take(3).sum::<u64>());

    return Ok(());
}


