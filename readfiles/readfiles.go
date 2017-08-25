// readfiles.go

package main



import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {
	const N_PROGRESS int = 10000
	buf := make([]byte, 65536)
	c := 0

	rd := bufio.NewReader(os.Stdin)
	for {

		fname, rd_err := rd.ReadString('\n')
		if rd_err != nil {
			break
		}
		fn := strings.TrimSpace(fname)

		f, e := os.Open(fn)
		if e != nil {
			log.Fatal(e)
		}
		defer f.Close()
		f.Read(buf)

		if (c % N_PROGRESS == 0) {
			fmt.Fprintf(os.Stderr, "\r%d", c)
		}
		c++
	}
	fmt.Fprintf(os.Stderr, "\r%d\n", c)
}

// EOF
