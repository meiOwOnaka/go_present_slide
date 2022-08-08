package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	now := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := now.In(jst)
	fmt.Println(nowJST.Format(time.RFC1123))
}
