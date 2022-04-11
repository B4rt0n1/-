package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)
	guessed := false
	for i := 9; i >= 0; i-- {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		inputnumber, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Входное значение должно содержать число")
			i++
			continue
		}
		if inputnumber > 100 || inputnumber < 1 {
			fmt.Println("Цифра не может быть больше 100 и меньше 1")
			i++
			continue
		}
		if inputnumber > n || inputnumber < n {
			fmt.Println("Oops. Your guess was", check(inputnumber, n))
			fmt.Println(strconv.Itoa(i), "attempts left")
			fmt.Println()
			continue
		}

		fmt.Println("Congratulations! You guessed the number")
		guessed = true
		break

	}
	if !guessed {
		fmt.Println("Sorry. You didn’t guess my number. It was:", n)
	}
}

func check(a int, b int) string {
	if a > b {
		return "HIGH"
	} else {
		return "LOW"
	}
}
