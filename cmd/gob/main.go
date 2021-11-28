package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/siuyin/enums-present/compass"
)

var network bytes.Buffer

func main() {
	enc := gob.NewEncoder(&network) // encoder to write to network

	d := compass.North
	if err := enc.Encode(d); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("On the wire as gob: %q\n", network.String())

	var dOut compass.Dir
	dec := gob.NewDecoder(&network) // decoder reads from the network
	if err := dec.Decode(&dOut); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("type: %T, value: %v, int-val:%d, string-val:%s\n", d, d, d, d)
}
