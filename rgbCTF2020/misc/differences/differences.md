## Description

If a difference could difference differences from differences to calculate differences, would the difference differently difference differences to difference the difference from different differences differencing?

## Included Files

DifferenceTest.java

## Writeup

Upon first opening the file it's easy to notice that some characters aren't correct. I opened it up in hexedit and found the first value and, taking from the heavy-handed hints in the description, I subtracted the hex value for 't' from it since that made sense for what would finish off "impor" and the resultant value was 0x72, or 'r'. Knowing that the flag begins with an r, it seemed like each of the corrupted characters had been incremented by the values of a flag character. So, I used a mixture of knowledge of Java with some common sense and corrected the file and then wrote a simple Go program that goes through each character and prints out all the differences that don't come out to 0.


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
