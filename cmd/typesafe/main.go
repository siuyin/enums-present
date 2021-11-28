package main

import (
	"fmt"

	"github.com/siuyin/enums-present/compass"
)

func main() {
	d := compass.North
	d = 2 // try d = int(2) and compass.Dir(2)
	fmt.Printf("type: %T, value: %v, int-val:%d, string-val:%s\n", d, d, d, d)
}
