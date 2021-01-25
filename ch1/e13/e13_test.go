// Result:

/* go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/nodeselector/tgp/ch1/e13
BenchmarkConcatEcho-16          1000000000               0.000030 ns/op
BenchmarkJoinEcho-16            1000000000               0.000064 ns/op
BenchmarkBufferEcho-16          1000000000               0.000032 ns/op
PASS
ok      github.com/nodeselector/tgp/ch1/e13     0.073s
package main */


// This is somewhat unexpected, since concat should take more time to run given that strings are immutable and must be copied

package main

import(
	"strings"
	"testing"
	"bytes"
	"os"
	"bufio"
)

func ReadInput() ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
        return nil, err
    }
    defer f.Close()

	var args []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	return args, scanner.Err()
}
func BenchmarkConcatEcho(b *testing.B) {
	args, _ := ReadInput()
	var s, sep string
	for i := range args {
		s += args[i]
		s += sep
	}
}
func BenchmarkJoinEcho(b *testing.B) {
	args, _ := ReadInput()
	echo := []string{}
	for _, arg := range args {
		echo = append(echo, arg)
	}
	strings.Join(echo, " ")
}


func BenchmarkBufferEcho(b *testing.B) {
	args, _ := ReadInput()
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(arg)
		buf.WriteString(" ")
	}
	buf.String()
}

