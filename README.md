TP0 – Programming Assignment in Go

This project is part of Trabajo Práctico 0, an introductory programming assignment written in Go.
The goal was to implement and test several basic functions, and then use them in a main program that processes data from text files.

📂 Project Structure
tp0/ (module root)
│── go.mod
│── main.go
│── ejercicios/
│     │── tp0.go
│     │── tp0_test.go


tp0.go → contains the implementation of the required functions.

tp0_test.go → contains the test cases (provided by the course).

main.go → main program that reads input files, compares arrays, sorts them, and prints the result.

📝 Implemented Functions

Inside tp0.go, the following functions were completed:

Swap() – swaps two elements.

Maximo() – returns the maximum value in an array.

Comparar() – compares two arrays.

Seleccion() – selection sort implementation.

Suma() – sums all elements in an array.

EsCadenaCapicua() – checks if a string is a palindrome.

▶️ How It Works

The program reads two input files: archivo1.in and archivo2.in.

Each file contains one integer per line.

Each file represents an array of integers.

It compares both arrays using the implemented functions (Comparar, Seleccion).

It prints the largest array, sorted in ascending order.

Output format: one number per line.

Example:
If the largest array is [2, 1, 7, 3], the output will be:

1
2
3
7

⚙️ Installation & Usage

Initialize the module:

go mod init tp0
go mod tidy


Run tests (inside ejercicios):

cd ejercicios
go test


Build the project (from module root):

go build tp0


Execute the program:

./tp0

✅ Testing

All functions are verified with the provided file tp0_test.go.
Run:

go test


All tests must pass before running the main program.

📌 Notes

Only tp0.go must be modified.

The provided tests (tp0_test.go) help ensure correctness.

Code must be formatted before submission:

go fmt tp0.go

🎯 Objective

This project serves as an introduction to Go modules, package organization, and unit testing.
It combines basic algorithmic functions with file handling and array processing in Go.
