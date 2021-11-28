package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/siuyin/enums-present/compass"
)

func main() {
	d := compass.North // try compass.North.String()
	b, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
