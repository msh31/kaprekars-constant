<div align="center">

# kaprekars-constant
Kaprekar's Constant implementation in Go
</div>

## What is kaprekar's constant?
6174, also knowon as kaprekar's constant, is a number that can be reached by repeatedly applying a specific process to any four-digit number that has at least two different digits. 
The process is as follows:
1. Take any four-digit number, using at least two different digits.
2. Arrange the digits in descending and then in ascending order to create two four-digit numbers, adding leading zeros if necessary.
3. Subtract the smaller number from the bigger number.
4. Go back to step 2 and repeat the process.
5. You will eventually reach the fixed point (6174) in at most 8 iterations.

This program computes the kaprekar's constant for any valid four digit number as described above and returns the number of steps it takes for the number to become a kaprekar's constant.
