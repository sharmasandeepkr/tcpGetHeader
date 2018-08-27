package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	services := os.Args[1]
	tcpAdd, err := net.ResolveTCPAddr("tcp", services)
	if err != nil {
		panic(err)
	}
	tcpCon, err := net.DialTCP("tcp", nil, tcpAdd)
	if err != nil {
		panic(err)
	}
	// var b []byte
	tcpCon.Write([]byte(" HEAD / HTTP/1.0\r\n\r\n"))
	// _, err = tcpCon.Read(b)

	res, err := ioutil.ReadAll(tcpCon)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
