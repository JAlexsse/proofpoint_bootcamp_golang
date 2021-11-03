package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	//three part loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//the while loop in go
	rand.Seed(time.Now().UnixNano()) //para alimentar nuestros numeros random
	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
	}
	fmt.Println("Got", i, "and broke loop.")

	//infinite loop
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go mylogger.ListerForLog(ch) //va a estar esperando input hasta que termine el programa

	fmt.Println("Enter something")
	for i := 0; i < 5; i++ {
		fmt.Println("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
