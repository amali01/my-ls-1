<h1 align=center>

![Push-swap](PushSwap.png)

</h1>


# Push-Swap Project

Push-Swap is a simple project that involves sorting an array of integers using two stacks and a set of instructions. This project consists of two programs:

1. `push-swap`: Calculates and displays the smallest program using the Push-Swap instruction language to sort the integer arguments received.
2. `checker`: Takes integer arguments and reads instructions from the standard input. It executes these instructions and displays "OK" if the integers are sorted, or "KO" if they are not.

## Table of Contents

- [Getting Started](#getting-started)
- [Usage](#usage)
- [Instructions](#instructions)
- [Testing](#testing)
- [Authors](#authors)

## Getting Started

To get started with the Push-Swap project, follow these steps:

1. Clone the repository to your local machine.

```console 
git clone https://learn.reboot01.com/git/amali/push-swap
```

2. Compile the `push-swap` and `checker` programs using the provided build file .
```console
sh build.sh
```
3. You can check the Result-log.txt.

## Usage

### `push-swap`

The `push-swap` program calculates and displays the smallest sequence of Push-Swap instructions to sort a given list of integer arguments. Here's how you can use it:

```console
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
$ ./push-swap "0 one 2 3"
Error
$ ./push-swap
$
```
### `checker`

The `checker` program reads a sequence of instructions from the standard input and executes them to verify if the given integers are sorted. Here's how you can use it:

```console
$ ./checker "3 2 1 0"
sa
rra
pb
KO
$ echo -e "rra\npb\nsa\n" | ./checker "3 2 one 0"
Error
$ echo -e "rra\npb\nsa\nrra\npa"
rra
pb
sa
rra
pa
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
$ ./checker
$
```
## Instructions

You can use the following instructions to manipulate the stacks and sort the integers:

- `pa`: Push the top element of stack B to stack A.
- `pb`: Push the top element of stack A to stack B.
- `sa`: Swap the first two elements of stack A.
- `sb`: Swap the first two elements of stack B.
- `ss`: Execute `sa` and `sb` simultaneously.
- `ra`: Rotate stack A (shift up all elements by 1; the first element becomes the last).
- `rb`: Rotate stack B.
- `rr`: Execute `ra` and `rb` simultaneously.
- `rra`: Reverse rotate stack A (shift down all elements by 1; the last element becomes the first).
- `rrb`: Reverse rotate stack B.
- `rrr`: Execute `rra` and `rrb` simultaneously.

## Testing

To test the push-swap and checker programs, you can create test cases with various input integers and instructions. Verify that the programs produce the expected output.

## Authors

- emahfood
- amali

 
## To be noted for the audits
```
ARG=("4 67 3 87 23"); ./push-swap "$ARG" 
```
$ARG should be in "" to be one argument 

- if you want to test a data of 100 int
```
79 1 28 14 27 16 78 84 11 18 69 30 42 7 100 52 56 73 54 25 43 80 32 77 4 5 68 98 81 72 62 70 53 95 34 15 29 91 67 59 2 48 99 71 23 35 64 76 31 87 47 22 75 90 6 24 13 92 83 45 97 86 46 19 66 89 50 74 58 20 44 40 88 39 0 65 33 36 8 63 3 60 49 37 21 10 17 55 38 9 94 96 85 82 57 12 26 51 61 41
```