package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"
)

func main() {
	s := "qwerty"
	context.WithTimeout(context.Background(), 5*time.Second)

	s1 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s1)

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
