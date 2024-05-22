# Math Skills

This repository contains a program written in Go for calculating various statistics such as Average, Median, Variance, and Standard Deviation from a given dataset.
Requirements

    Go (latest version)

## Installation

Make sure you have the latest version of Go installed on your system.
You can download and install it from the official [Go website](https://go.dev/dl/).

## Usage

Clone this repository to your local machine:
```bash
git clone https://learn.zone01kisumu.ke/git/hilaokello/math-skills
```
Navigate to the project directory:

```bash
cd math-skills
```
Run the program with the following command:
```bash
go run main.go data.txt
```
Replace data.txt with the path to your dataset file. The dataset file should contain one value per line, representing a statistical population.
Example

If your dataset file data.txt contains:
```txt
189
113
121
114
145
110
...
```
Running the program will output:

```bash
Average: 128
Median: 118
Variance: 1115
Standard Deviation: 33
```
The values are rounded to the nearest integer.

## Structure of the Program

The **main.go** file contains the main program logic. It reads data from a file, calculates the statistics, and prints the results.

The **formulas** package contains functions for calculating the statistics mentioned.

The **formulas_test.go** file contains the tests to all my functions.

## Author

[Hilary Okello](https://github.com/HilaryOkello)

## License

This project is licensed under the [MIT License](./LICENSE.txt).