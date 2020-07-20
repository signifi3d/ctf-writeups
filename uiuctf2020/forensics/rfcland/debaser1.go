package main

import (
	b64 "encoding/base64"
	"os"
	"bufio"
	"fmt"
)

func main() {
	fd, _ := os.Open("data_packets.pcap")
	new_fd, _ := os.Create("./decoded")

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	
	for scanner.Scan() { 
		eq_count := 0
		scan_str := scanner.Text()[0:len(scanner.Text())]
		in_str, scan_err := b64.StdEncoding.DecodeString(scan_str)

		for scan_err != nil && eq_count < 2 {
			scan_str = scan_str + "="
			eq_count++
			in_str, scan_err = b64.StdEncoding.DecodeString(scan_str)
		}

		new_fd.Write(in_str)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return

}
