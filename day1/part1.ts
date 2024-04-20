import { getData, lines, isNumeric } from "../lib/data";

getData("https://adventofcode.com/2023/day/1/input").then((data) => {
    console.log(lines(data).reduce((accumulator, line) => accumulator + solveByLine(line), 0));
});

function solveByLine(line: string): number {
    let count = "";

    for (let i = 0; i < line.length; i++) {
        if (isNumeric(line[i])) {
            count = count + line[i];
            break; // Exit the loop once the target is found
        }
    }

    for (let i = line.length - 1; i >= 0; i--) {
        if (isNumeric(line[i])) {
            count = count + line[i];
            break; // Exit the loop once the target is found
        }
    }

    return Number(count)
}
