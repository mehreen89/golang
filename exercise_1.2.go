package main
import (
	"fmt"
	"os"
)
func main() {
	for V, arg := range os.Args[1:] {
		fmt.Println("Index",V ,"contains the value ",arg )
	}
}
