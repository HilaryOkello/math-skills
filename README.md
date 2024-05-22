# Math Skills

This repository contains a program written in Go for calculating various statistics such as Average, Median, Variance, and Standard Deviation from a given dataset.
### Requirements
You need to have at least  go version go1.22.0 or higher to run the program. You can download and install it from the official [Go website](https://go.dev/dl/).

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
or 
```bash
go run . data.txt
```

Replace data.txt with your dataset file (Must be a text file with a ".txt" extension). The dataset file should contain one value per line, representing a statistical population.

### Example

If your dataset file data.txt contains:
```txt
189
113
121
114
145
110
```
Running the program will output:

```bash
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
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