## Description

java SecureRandom is supposed to be, well, secure, right? nc challenge.rgbsec.xyz 7425

## Included files

Main.java

## Writeup

Browsing through the source code, the general idea of the program is that it seeds SecureRandom with a randomly determined time interval, presents you the first five numbers it generates, then it generates the 5 numbers you have to guess, and you get 20 tries to guess it. Since it's only 5 numbers that can only be 1-5 there are only 3125 possbilities. With 20 tries it doesn't seem too crazy to just generate your own random possibilities and automate the guessing. Here is my solution that does that written in Go.

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
