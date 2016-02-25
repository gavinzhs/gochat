package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryCoding(t *testing.T) {
	data := []byte("hello")
	bf := bytes.NewBuffer(nil)
	binary.Write(bf, binary.BigEndian, data)
	size := binary.Size(data)
	o := make([]byte, size)
	binary.Read(bf, binary.BigEndian, o)
	fmt.Printf("这是原始数据:%s\n", o)
}

func TestBinaryCoding1(t *testing.T) { //这里是uint32
	data := uint32(100000)
	bf := bytes.NewBuffer(nil)
	binary.Write(bf, binary.BigEndian, data)
	size := binary.Size(data)
	o := make([]byte, size)
	binary.Read(bf, binary.BigEndian, o)
	fmt.Printf("这是原始数据:%d, 大小: %d\n", o, size)
}
