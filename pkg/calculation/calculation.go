package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func indexOf(slice []string, values string) int {
	for i, v := range slice {
		if strings.Contains(values, v) {
			return i
		}
	}
	return -1
}

func isValidInNumber(num string, c rune) bool {
	validChars := "0123456789"
	if c == '.' {
		return !strings.Contains(num, ".") && len(num) != 0
	}
	if c == '-' {
		return !strings.Contains(num, "-") && len(num) == 0
	}
	return strings.Contains(validChars, string(c))
}

func isValidOperation(c rune) bool {
	validChars := "+-*/"
	return strings.Contains(validChars, string(c))
}

func Calc(expression string) (float64, error) {
	var err error
	expression = strings.ReplaceAll(expression, " ", "")
	expression, err = evaluate(expression)
	if err != nil {
		return 0.0, err
	}
	s := []string{""}
	for _, i := range expression {
		if isValidInNumber(s[len(s)-1], i) {
			s[len(s)-1] += string(i)
		} else if isValidOperation(i) {
			s = append(s, string(i))
			s = append(s, "")
		} else {
			return 0, errors.New("Expression is not valid")
		}
	}
	var ind int
	var n, n1, n2 float64
	for len(s) != 1 {
		if ind = indexOf(s, "*/"); ind != -1 {
			n1, err = strconv.ParseFloat(s[ind-1], 64)
			if err != nil {
				return 0, err
			}
			n2, err = strconv.ParseFloat(s[ind+1], 64)
			if err != nil {
				return 0, err
			}
			if s[ind] == "*" {
				n = n1 * n2
			} else if s[ind] == "/" {
				if n2 == 0 {
					return 0, errors.New("Expression is not valid")
				}
				n = n1 / n2
			}
		} else if ind = indexOf(s, "+-"); ind != -1 {
			n1, err = strconv.ParseFloat(s[ind-1], 64)
			if err != nil {
				return 0, err
			}
			n2, err = strconv.ParseFloat(s[ind+1], 64)
			if err != nil {
				return 0, err
			}
			if s[ind] == "+" {
				n = n1 + n2
			} else if s[ind] == "-" {
				n = n1 - n2
			}
		}
		s[ind+1] = strconv.FormatFloat(n, 'f', -1, 64)
		s = append(s[:ind-1], s[ind+1:]...)
	}
	n, err = strconv.ParseFloat(s[0], 64)
	return n, err
}

func evaluate(expression string) (string, error) {
	s := []string{""}
	var n int
	var v float64
	var err error
	for _, i := range expression {
		if i == '(' {
			n++
			if n == 1 {
				s = append(s, "")
			} else {
				s[len(s)-1] += string(i)
			}
		} else if i == ')' {
			n--
			if n == 0 {
				v, err = Calc(s[len(s)-1])
				if err != nil {
					return "", err
				} else {
					s[len(s)-1] = fmt.Sprintf("%v", v)
					s = append(s, "")
				}
			} else if n < 0 {
				return "", errors.New("Expression is not valid")
			} else {
				s[len(s)-1] += string(i)
			}
		} else {
			s[len(s)-1] += string(i)
		}
	}
	if n == 0 {
		return strings.Join(s, ""), nil
	} else {
		return "", errors.New("Expression is not valid")
	}
}
