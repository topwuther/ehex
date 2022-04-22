package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

var name = flag.String("name", "", "input file name")
var offset = flag.Int64("offset", 0x0, "input offset")
var data = flag.Int64("data", 0, "input data")

func main() {
	flag.Parse()
	file, _ := os.OpenFile(*name, os.O_RDWR, 0644)
	buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.BigEndian, uint8(*data))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = file.WriteAt(buffer.Bytes(), *offset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
