package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	fmt.Println("start Write")
	_, err = conn.Write(msg[0:len])
	checkError(err)

	fmt.Println("start Read")
	_, err = conn.Read(msg[0:])
	fmt.Println("end read")
	checkError(err)

	fmt.Println("Got responce")
	if msg[5] == 13 {
		fmt.Println("Indentifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matchees")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0
	n := 0
	for n+1 < len(msg) {
		sum += (int(msg[n] << 8)) | int(msg[n+1])
		n++
	}
	if n < len(msg) {
		sum += (int(msg[n]) << 8)
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)

}

func checkError(err error) {
	fmt.Println("checkError");
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
