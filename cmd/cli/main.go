package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := flag.String("c", "", "Math equation")
	flag.Parse()

	if *input == "" {
		*input, _ = readStdin()
	}

	equation := strings.Join(strings.Fields(*input), "")

	operator := regexp.MustCompile(`[+*/-]`).FindAllString(equation, -1)[0]
	s := regexp.MustCompile(`[+*/-]`).Split(equation, -1)
	if len(s) < 2 {

	}
	first, _ := strconv.Atoi(s[0])
	second, _ := strconv.Atoi(s[1])

	switch operator {
	case "+":
		result := first + second
		fmt.Printf("Result: %d\n", result)
	case "-":
		result := first - second
		fmt.Printf("Result: %d\n", result)
	case "*":
		result := first * second
		fmt.Printf("Result: %d\n", result)
	case "/":
		result := first / second
		fmt.Printf("Result: %d\n", result)
	}

}

func readStdin() (string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return "", errors.New("StdIn not a named pipe")
	}

	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
