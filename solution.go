package main

import (
	"bufio"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	/* Создать массив на 10 GB */
	ballast := make([]byte, 10<<30)

	/* Открыть файл */
	file, err := os.Open("ip_addresses")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	/* Читать файл построчно */
	in := bufio.NewReader(file)
	var i int64

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		ip := net.ParseIP(line)
		if ip == nil {
			fmt.Println("Invalid IP address")
		}
		i = IP4toInt(ip)
		ballast[i] = 1
	}

	/* Подсчет уникальных значений */
	result := 0
	for i := range ballast {
		if ballast[i] == 1 {
			result++
		}
	}
	fmt.Println("result:", result)
	fmt.Println("Execution time:", time.Since(start))

	// [Output] result: 1000000000
	// Execution time: 16m45s
}

// IP4toInt : Конвертировать Ip-адрес в число
func IP4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}
