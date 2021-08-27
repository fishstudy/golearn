package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":8888", nil))
	}()
	cidrs := []string{
		"1.2.3.4/8",
		"1.2.3.4/16",
		"1.2.3.4/24",
		"1.2.3.4/32",
		"1.2.3.4/1",
		"1.2.3.4/7",
		"1.2.3.4/22",
	}
	ipInt := [][2]int{
		{16777216, 33554431},
		{16908288, 16973823},
		{16909056, 16909311},
		{16909060, 16909060},
		{0, 2147483647},
		{0, 33554431},
		{16908288, 16909311},
	}
	for {
		for times := 0; times < 1000; times++ {
			for index := 0; index < len(cidrs); index++ {
				cidr := cidrs[index]
				targetStart, targetEnd := uint32(ipInt[index][0]), uint32(ipInt[index][1])
				start, end, _ := Cidr2range(cidr)
				if start != targetStart || end != targetEnd {
					fmt.Printf("cidr:%s, targetStart:%d, targetEnd:%d, start:%d, end:%d\n", cidr, targetStart, targetEnd, start, end)
				}
			}
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

// Ip2uint 将ip转换成unit32格式
func Ip2uint(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// Cidr2range 将cidr格式的ipv4地址转换成[uint32 startIp, uint32 endIp]形式
func Cidr2range(cidr string) (uint32, uint32, error) {
	vipSegs := strings.Split(cidr, "/")
	ip := Ip2uint(vipSegs[0])

	var startIp uint32
	var endIp uint32

	if len(vipSegs) > 1 {
		mask, err := strconv.Atoi(vipSegs[1])
		if err != nil {
			return 0, 0, fmt.Errorf("mask type conversion failed")
		}
		startIp = ip & (((1 << uint32(mask)) - 1) << uint32(32-mask))
		endIp = ip | ((1 << uint32(32-mask)) - 1)

	} else {
		startIp = ip
		endIp = ip
	}
	return startIp, endIp, nil
}