package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]

	if !isNumeric(arg1) || !isNumeric(arg3) {
		return
	}
	res := 0
	switch arg2 {
	case "+":
		res = Atoi(arg1) + Atoi(arg3)
	case "-":
		res = Atoi(arg1) - Atoi(arg3)
	case "/":
		if arg3 == "0" {
			os.Stdout.WriteString("No division by 0")
			return
		}
		res = Atoi(arg1) / Atoi(arg3)
	case "%":
		if arg3 == "0" {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		res = Atoi(arg1) / Atoi(arg3)

	case "*":
		res = Atoi(arg1) * Atoi(arg3)

	}
	if atoi >  
	os.Stdout.WriteString(itoa(res))
}

func Atoi(s string) int {
	result := 0
	t := 1
	for i, c := range s {
		if (c == '-' || c == '+') && i == 0 {
			if c == '-' {
				t = -1
			}
			continue
		} else if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result * t
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func itoa(n int) string {
	str := ""
	for n >= 0 {
		str = string(n%10 + '0')
		n /= 10
	}
	return str
}
