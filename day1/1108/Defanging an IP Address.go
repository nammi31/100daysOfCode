
package main
import (
	"fmt"
	"strings"
)
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]",len(address)/2)
}
func main( ){
	address:="255.100.50.0"
	fmt.Println(defangIPaddr(address))
}