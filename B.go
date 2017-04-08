package main

import (
       "fmt"
       "bufio"
       "os"
       "strconv"
       "strings"
)

func main() {
     reader := bufio.NewReader(os.Stdin)
     strcases,_ := reader.ReadString('\n')
     cases,_ := strconv.Atoi(strings.Trim(strcases,"\n"))
     for i:=1; i<=cases; i++ {
     	input, _ := reader.ReadString('\n')
		input = strings.Trim(input,"\n")
		N, _ := strconv.Atoi(input)
		found := false
		for found==false {
			copyN := N
			if N <= 0 {
				break
			}
			//set prevDigit to highest digit for first pass
			prevDigit := 9
			tidy := true
	 		for copyN>0 {
	    		digit := copyN%10
	    		if digit > prevDigit {
	    			tidy = false
					break;
	    		}	     
	    		copyN = copyN/10
	    		prevDigit = digit
	 		}
	 		if !tidy {
		   		N = N-1
	 		} else {
	 			found = true
	 		}
		}	
		fmt.Printf("Case #%d: %d\n",i,N)
    }
}