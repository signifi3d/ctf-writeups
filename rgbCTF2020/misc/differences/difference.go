package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	orig, err1 := ioutil.ReadFile("DifferenceTest.java")
	corr, err2 := ioutil.ReadFile("corrected.java")
	if err1 != nil || err2 != nil {
		fmt.Printf("Error opening files.\n")
		return
	}

	for i, curr := range orig {
		diff := curr - corr[i]
		if diff != 0 {
			fmt.Printf("%c", diff)
		}
	}

	return
}
