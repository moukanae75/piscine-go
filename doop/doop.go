package main

import (
	"os"
)

const (
	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, err1 := Atoi(os.Args[1])
	op := os.Args[2]
	num2, err2 := Atoi(os.Args[3])

	if !err1 || !err2 {
		return
	}

	switch op {
	case "+":

		if (num1 > 0 && num2 > 0 && num1 > MaxInt64-num2) ||
			(num1 < 0 && num2 < 0 && num1 < MinInt64-num2) {
			return
		}
		PrintNbr(int(num1) + int(num2))

	case "-":

		if (num1 > 0 && num2 < 0 && num1 > MaxInt64+num2) ||
			(num1 < 0 && num2 > 0 && num1 < MinInt64+num2) {
			return
		}
		PrintNbr(int(num1) - int(num2))
	case "*":

		if num1 == 0 || num2 == 0 {
			os.Stdout.WriteString("0")
			return
		}

		maxAllowed := MaxInt64 / abs(int64(num1))
		minAllowed := MinInt64 / abs(int64(num1))

		if (num1 > 0 && num2 > 0 && num2 > int(maxAllowed)) ||
			(num1 < 0 && num2 < 0 && num2 < int(minAllowed)) ||
			(num1 > 0 && num2 < 0 && num2 < int(minAllowed)) ||
			(num1 < 0 && num2 > 0 && num2 > int(maxAllowed)) {
			return
		}
		PrintNbr(int(num1) * int(num2))

	case "/":

		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}

		if num1 < num2 {
			os.Stdout.WriteString("0\n")
			return
		}

		if num1 == MinInt64 && num2 == -1 {
			return
		}
		PrintNbr(int(num1) / int(num2))

	case "%":

		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}

		if num1 == MinInt64 && num2 == -1 {
			return
		}
		PrintNbr(int(num1) % int(num2))

	default:
		return
	}
	os.Stdout.WriteString("\n")
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func PrintNbr(n int) {
	str := ""
	last := ""
	if n < 0 {
		n = -n
		last = "-"

	}
	for n > 0 {
		str = string(n%10+48) + str
		n /= 10
	}
	last += str
	os.Stdout.WriteString(last)
}

func Atoi(s string) (int, bool) {
	nbr := 0
	if len(s) > 0 {
		factor := 1
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] < '0' || s[i] > '9' {
				if i != 0 || (s[0] != '-' && s[0] != '+') {
					return 0, false
				}
			}
			if s[i] != '-' && s[i] != '+' {
				nbr += (int(s[i]) - 48) * factor
				factor = factor * 10
			}
		}
		if s[0] == '-' {
			return -nbr, true
		}
	}
	return nbr, true
}
