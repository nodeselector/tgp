package main

import(
	"fmt"
	"os"
	"bytes"
)

func main() {
	var buf bytes.Buffer
	for _, arg := range os.Args {
		buf.WriteString(arg)
		buf.WriteString(" ")
	}
	fmt.Println(buf.String())
}