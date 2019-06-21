package main

import "fmt"
import "math/rand"
import "github.com/Infoblox-CTO/golangtraining/firstlib"

func main() {
	fmt.Println("hello world")
	tryit(100)

}

func tryit(limit int){
	thenumber := rand.Intn(limit) + 1
	var guess int
	for i := 0; i < 20; {
		fmt.Print("Your guess? ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Try again.")
			i++
			continue
		}
		message, done := firstlib.Checkguess(guess, thenumber)
		fmt.Println(message)
		if done {
			break
		}
		i++
	}
}