# Guess-it-1
This project involves building a program that takes a sequence of numbers as input and predicts a range for the next number in the sequence.

## Features
1. **[Mean](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/average.go)** This is function that calculates the average of the population.
2. **[Median](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/median.go):** It is the middle value in a set of numbers when they are arranged in order from smallest to largest
3. **[Variance](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/variance.go):** This function checks the variance of the population and calculates the variance of the data using the mean and the data itself.
4. **[Standard Deviation](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/standarddeviation.go):** This function uses the variance function to compute the standard deviation of the population data.
## Cloning
To clone this repository to see how the program works use the command below on your bash terminal;
```bash
git clone https://github.com/Wendy-Tabitha/guess--it--1.git
```


## Instructions
The program reads a number from standard input and prints a range for the next expected number. The range should be printed in the format lower_limit upper_limit.

## Run command
Before running this command below change the directory to the working directory by using this command below;
```bash
cd guess--it--1
```
To run the program copy this to your terminal;
```bash
go run .
```

## Implementation Details

- Reading Input: The program should continuously read numbers from standard input.
- Calculating Ranges: In this program there is use of the function for mean and standard deviation which are used in finding the range where the next number is.
- Outputting the Range: For each input number, output the predicts the range for the next number.

## Testing

This project has test files which test the fuctionality of the functions that calculates the average, median, variance and standard deviation. to run the test file:

```bash
go test -v
```
We run the above command in the console or terminal to test the files that are in the mathskills directory.

## Learning Objectives

This project is designed to help learn about:

- Statistics and Mathematics.
- Method that predicts a range for the next number in the sequence.
- Average.
- Median.
- Variance.
- Standard Deviation.

This project was done by:
[Wendy-Tabitha](https://github.com/Wendy-Tabitha/guess--it--1)

## License
This project is licensed under the terms of the [MIT License](./LICENSE).