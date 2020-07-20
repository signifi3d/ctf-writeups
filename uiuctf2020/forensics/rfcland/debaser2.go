package main

import (
	b64 "encoding/base64"
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	fd, _ := os.Open("data_packets.pcap")
	beg_count := 0
	mid_count := 0
	end_count := 0

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	
	for scanner.Scan() { 
		eq_count := 0
		scan_str := scanner.Text()[0:len(scanner.Text())]
		in_str, scan_err := b64.StdEncoding.DecodeString(scan_str)
		new_fd, _ := os.Create("./holder")

		for scan_err != nil && eq_count < 2 {
			scan_str = scan_str + "="
			eq_count++
			in_str, scan_err = b64.StdEncoding.DecodeString(scan_str)
		}

		if in_str[0] == '\xff' && in_str[1] == '\xd8' {
			new_fd, _ = os.Create("./begins/beg" + strconv.Itoa(beg_count))
			beg_count++
		} else if in_str[len(in_str)-2] == '\xff' && in_str[len(in_str)-1] == '\xd9' {
			new_fd, _ = os.Create("./ends/end" + strconv.Itoa(end_count))
			end_count++
		} else {
			new_fd, _ = os.Create("./mids/mid" + strconv.Itoa(mid_count))
			mid_count++
		}

		new_fd.Write(in_str)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return

}
