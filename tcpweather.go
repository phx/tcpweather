package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func doDial(host, port, cmd string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection established between %s and localhost.\n", host)
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	var buf = make([]byte, 4096)
	len, err := conn.Read(buf)
	if err != nil {
		log.Panicln(err)
	}
	msg := string(buf[:len])
	if strings.Contains(msg, "Press Return to continue") {
		fmt.Println(msg)
		conn.Write([]byte("\n"))
		len, err = conn.Read(buf)
		if err != nil {
			log.Panicln(err)
		}
		msg = string(buf[:len])
		fmt.Println(msg)
		conn.Write([]byte("CHI\n"))
		len, _ = conn.Read(buf)
		msg = string(buf[:len])
		fmt.Println(msg)
	}
}

func main() {
	cmd := "geoip"
	host := "rainmaker.wunderground.com"
	port := "23"
	doDial(host, port, cmd)
}
