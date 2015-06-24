package main
import (
	"github.com/ohisama/mci"
	"fmt"
)
func main() {
	var cmd string
	cmd = "open zashop.mid"
	mci.Send(cmd)
	cmd = "play zashop.mid notify"
	mci.Send(cmd)
	var input string
	fmt.Scanln(&input)
}
