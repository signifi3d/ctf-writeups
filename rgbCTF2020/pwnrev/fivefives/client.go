package main

import (
	"strings"
	"fmt"
	"net"
	"bufio"
	"strconv"
	"math/rand"
	"time"
)

//Seed doesn't really matter that much
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//Generates guesses
func getRandString() string {
	ret := ""

	for i := 0; i < 5; i++ {
		ret += strconv.Itoa(r.Intn(5)+1) + " "
	}

	ret += "\n"

	return ret
}

func main() {
	//Just keep running it until correctly guessed
	for {
		//Establish connection
		serv := "167.172.123.213:7425"
		conn, err := net.Dial("tcp", serv)
		if err != nil {
			fmt.Println(err)
			return
		}
		reader := bufio.NewReader(conn)
	
		//Take in the first few messages from the server
		message, _ := reader.ReadString('\n')
		fmt.Print(message)
		message, _ = reader.ReadString('\n')
		fmt.Print(message)
		message, _ = reader.ReadString('\n')
		fmt.Print(message)
		message, _ = reader.ReadString('\n')
		fmt.Print(message)
		message, _ = reader.ReadString('\n')
		fmt.Print(message)
		
		//Start the 20 guesses, extra print is just because I like to see it
		fmt.Fprintf(conn, "20\n")
		fmt.Print("20\n")
		
		//Go through the 20 guesses
		for i:=0; i < 20; i++ {
			message, _ = reader.ReadString('\n')
			fmt.Print(message)
			randStr := getRandString()
			fmt.Print(randStr)
			fmt.Fprintf(conn, randStr)
			message, _ = reader.ReadString('\n')
			fmt.Print(message)
			//until flag is delivered
			if strings.Contains(message, "flag") {
				message, _ = reader.ReadString('\n')
				fmt.Print(message)
				return
			}
		}
	}
	return
}
