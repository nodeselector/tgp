package main

import(
	"fmt"
	"os"
	"bytes"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	for index, arg := range os.Args {
		buf.WriteString(strconv.Itoa(index))
		buf.WriteString(": ")
		buf.WriteString(arg)
		buf.WriteString("\n")
	}
	fmt.Println(buf.String())
}