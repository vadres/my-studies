package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var u int64 = int64(time.Now().Nanosecond())
	rand.Seed(u)
	
	fmt.Println(time.Now())
	
	fmt.Println("My favorite number is", rand.Intn(10))
}
