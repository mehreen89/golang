package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	
	var sum int
	for i:=1; i < len(os.Args); i++ {
		str:= os.Args[i]
		num, err := strconv.Atoi(str)
	    if err == nil {
	    	sum = sum + num
	    }
		
	}
	fmt.Println(sum)
}
