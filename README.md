<div align="center">

# kaprekars-constant
Kaprekar's Constant implementation in Go
</div>

## What is kaprekar's constant?
6174, also known as kaprekar's constant, is a number that can be reached by repeatedly applying a specific process to any four-digit number that has at least two different digits. 
The process is as follows:
1. Take any four-digit number, using at least two different digits.
2. Arrange the digits in descending and then in ascending order to create two four-digit numbers, adding leading zeros if necessary.
3. Subtract the smaller number from the bigger number.
4. Go back to step 2 and repeat the process.
5. You will eventually reach the fixed point (6174) in at most 8 iterations.

This program computes the kaprekar's constant for any valid four digit number as described above and returns the number of steps it takes for the number to become a kaprekar's constant.


## Program logic explained
First, we declare a 'main package' (this is just a group of all included files with that package name) then we import 2 packages, fmt for formatted text and printing to the console and sort to sort integers.

We then have a boolean function that takes an integer as paramater and returns whether the input is higher or equal to 1000 and lower or equal to 9998, true if that's the case, otherwise return false.

Then we have the main function, here we declare 3 integer variables 'n' 'count' and 'result', we clear the console screen using ascii code, then a for loop that asks the user for a 4 digit number, declares 2 variables '_' and err and assign a scan to it that returns the output to the n varialbe from earlier, if the error isn't nil then the user entered a letter instead of a number which will fail and the looop restarts, then if they pass that the number needs to be 4 digits, only then will the main loop start for kaprekar's constant.

Then we have a for loop that doesnt end unti the result is 6174, the fixed point. We increase the count each iteration, we assign the variable 'digits' to an arraay of integers that take the input (n) and divide it by 1000, divide that by 100 and then apply a modulo factor of 10, then take that numbber, repeat the process but divide by 10 this time and finally apply another factor 10 modulo to the number, then we use the imported sort package to sort the digits, then have a variable called 'largest' with the sorted array gets resorted to get the largest numberr (not quite sure how, takes the 3rd element and does it * 1000 then, 100 for the 2nd element, 10 for the 1st and lastly adds them all up in between that and lastly adds up the 1st digit?*) and the other way around for the smallest one, then we give the result variable the result of largest - ssmallest, give n the new value of result and print the largest, smallest and result and last but not least, we print in how many iterations we got to the fixed point of 6174.