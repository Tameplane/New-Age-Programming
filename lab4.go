package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number cannot have a real square root")
	}
	return math.Sqrt(n), nil
}

func ParseInput(input string) (float64, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}
	return value, nil
}

func testDivide() bool {
	result, err := Divide(10, 2)
	return err == nil && result == 5
}

func testDivideByZero() bool {
	_, err := Divide(10, 0)
	return err != nil
}

func testSquareRoot() bool {
	result, err := SquareRoot(25)
	return err == nil && result == 5
}

func testNegativeSquareRoot() bool {
	_, err := SquareRoot(-9)
	return err != nil
}

func testEmptyInput() bool {
	_, err := ParseInput("")
	return err != nil
}

func testInvalidInput() bool {
	_, err := ParseInput("abc")
	return err != nil
}

func testValidInput() bool {
	result, err := ParseInput("42")
	return err == nil && result == 42
}

func testSquareRootAccuracy() bool {
	result, err := SquareRoot(2)
	return err == nil && math.Abs(result*result-2) < 0.000001
}

func runTest(name string, passed bool) {
	if passed {
		fmt.Println("PASS:", name)
	} else {
		fmt.Println("FAIL:", name)
	}
}

func main() {
	fmt.Println("ERROR HANDLING DEMONSTRATION")

	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	result, err = SquareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}

	_, err = SquareRoot(-9)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = ParseInput("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = ParseInput("abc")
	if err != nil {
		fmt.Println("Error:", err)
	}

	result, err = ParseInput("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed input:", result)
	}

	fmt.Println("\nUNIT TEST RESULTS")

	runTest("TestDivide", testDivide())
	runTest("TestDivideByZero", testDivideByZero())
	runTest("TestSquareRoot", testSquareRoot())
	runTest("TestNegativeSquareRoot", testNegativeSquareRoot())
	runTest("TestEmptyInput", testEmptyInput())
	runTest("TestInvalidInput", testInvalidInput())
	runTest("TestValidInput", testValidInput())
	runTest("TestSquareRootAccuracy", testSquareRootAccuracy())
}
