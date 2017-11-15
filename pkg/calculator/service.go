package calculator

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func Calculate (equation string) string{
	fmt.Println(equation)

	equation = strings.Join(strings.Fields(equation),"")
	
	operator := regexp.MustCompile(`[+*/-]`).FindAllString(equation, -1)[0]
	s := regexp.MustCompile(`[+*/-]`).Split(equation, -1)
	if len(s) < 2 {
		
	}
	first, _ := strconv.Atoi(s[0])
	second, _ := strconv.Atoi(s[1])
	result := 0
	switch operator {
	case "+":
		result = first+second
		fmt.Printf("Result: %d\n", result)
	case "-":
		result = first-second
		fmt.Printf("Result: %d\n", result)
	case "*":
		result = first*second
		fmt.Printf("Result: %d\n", result)
	case "/":
		result = first/second
		fmt.Printf("Result: %d\n", result)
	}

	out := strconv.Itoa(result)

	fmt.Println(out)
	return out
}