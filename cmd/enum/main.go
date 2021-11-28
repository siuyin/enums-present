package main

import (
	"fmt"

	"github.com/siuyin/enums-present/compass"
)

func main() {
	fmt.Println("enums in Go")
	d := compass.East
	fmt.Printf("type: %T, value: %v, int-val:%d, string-val:%s\n", d, d, d, d)
}
