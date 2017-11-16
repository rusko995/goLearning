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

	equation := strings.Join(strings.Fields(*input), "") //delete empty space

	result, err := simpleCalc(equation) //for calculating simple equation without brackets and with 2 numbers
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result: %d\n", result)

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

func calculate(equation string) (int, error) {
	result := 0
	var err error

	if string(equation[0]) == "(" { //find matching bracket, erase it and pass this as left side
		openBracket := 1
		leftSide := 0
		rightSide := 0
		operator := ""

		for ix, char := range equation { //find matching brackets
			if string(char) == "(" {
				openBracket++
			} else if string(char) == ")" {
				openBracket--
			}

			if openBracket == 0 {
				leftSide, err = calculate(equation[1 : ix-1])
				if err != nil {
					return 0, err
				}
				if len(equation)-1 > ix {
					operator = string(equation[ix+1])
					rightSide, err = calculate(equation[(ix + 2):])
					if err != nil {
						return 0, err
					}
				}

				switch operator {
				case "+":
					result := leftSide + rightSide
					return result, nil
				case "-":
					result := leftSide - rightSide
					return result, nil
				case "*":
					result := leftSide * rightSide
					return result, nil
				case "/":
					result := leftSide / rightSide
					return result, nil
				}
			}

		}
	} else {
		out, err := strconv.Atoi(equation)
		if err != nil {
			return out, nil
		}

		//separate on + or -

	}
	return result, errors.New("You didn't pass input correctly")
}

func simpleCalc(equation string) (int, error) {
	var (
		operator       string
		prexifForFirst string
	)

	operatorExist := regexp.MustCompile(`[+*/-]`).FindAllString(equation, -1) //see if there is at least one operator

	if len(operatorExist) < 1 {
		return 0, errors.New("You didn't pass input correctly")
	} else if len(operatorExist) == 1 {
		operator = regexp.MustCompile(`[+*/-]`).FindAllString(equation, -1)[0]
	} else if len(operatorExist) == 2 {
		operatorSplit := regexp.MustCompile(`[+*/-]`).FindAllString(equation, -1)
		prexifForFirst = operatorSplit[0]
		operator = operatorSplit[1]
	}

	s := regexp.MustCompile(`[+*/-]`).Split(equation, -1)
	if len(s) < 2 {
		return 0, errors.New("You didn't pass input correctly")
	}

	first, _ := strconv.Atoi(s[0])
	second, _ := strconv.Atoi(s[1])

	if strings.Compare(prexifForFirst, "-") == 0 {
		first = first * (-1)
	}

	switch operator {
	case "+":
		result := first + second
		return result, nil
	case "-":
		result := first - second
		return result, nil
	case "*":
		result := first * second
		return result, nil
	case "/":
		result := first / second
		return result, nil
	}

	return 0, errors.New("Operator input error")
}
