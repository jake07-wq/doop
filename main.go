package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	val1, ok1 := atoi(args[0])
	val2, ok2 := atoi(args[2])
	op := args[1]

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		res := val1 + val2
		if (val2 > 0 && val1 > 9223372036854775807-val2) || (val2 < 0 && val1 < -9223372036854775808-val2) {
			return
		}
		printInt(res)
	case "-":
		res := val1 - val2
		if (val2 < 0 && val1 > 9223372036854775807+val2) || (val2 > 0 && val1 < -9223372036854775808+val2) {
			return
		}
		printInt(res)
	case "*":
		if val1 == 0 || val2 == 0 {
			printInt(0)
			return
		}
		res := val1 * val2
		if res/val1 != val2 {
			return
		}
		printInt(res)
	case "/":
		if val2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		printInt(val1 / val2)
	case "%":
		if val2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		printInt(val1 % val2)
	default:
		return
	}
}

func atoi(s string) (int64, bool) {
	var res int64
	sign := int64(1)
	i := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')
		// Overflow check during parsing
		if sign == 1 && (res > (9223372036854775807-digit)/10) {
			return 0, false
		}
		if sign == -1 && (-res < (-9223372036854775808+digit)/10) {
			return 0, false
		}
		res = res*10 + digit
	}
	return res * sign, true
}

func printInt(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	if n == -9223372036854775808 {
		os.Stdout.WriteString("-9223372036854775808\n")
		return
	}
	var res []byte
	sign := false
	if n < 0 {
		sign = true
		n = -n
	}
	for n > 0 {
		res = append([]byte{byte(n%10 + '0')}, res...)
		n /= 10
	}
	if sign {
		res = append([]byte{'-'}, res...)
	}
	res = append(res, '\n')
	os.Stdout.Write(res)
}
