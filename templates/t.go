package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
)

/* begin of MyIO */
const eof = 0

type any = interface{}

type MyIO struct {
	in  *os.File
	out *bufio.Writer
	_i  int
	_n  int
	buf []byte
}

func NewIO(in *os.File, out *os.File) *MyIO {
	return &MyIO{
		in,
		bufio.NewWriter(out),
		0, 0, make([]byte, 1<<12),
	}
}

func (myio *MyIO) rc() byte {
	if myio._i == myio._n {
		myio._n, _ = myio.in.Read(myio.buf)
		if myio._n == 0 { // EOF
			return eof
		}
		myio._i = 0
	}
	b := myio.buf[myio._i]
	myio._i++
	return b
}

func (myio *MyIO) ReadInt() (x int) {
	rc := myio.rc
	b := rc()
	neg := false
	for ; '0' > b || b > '9'; b = rc() {
		if b == eof {
			return
		}
		if b == '-' {
			neg = true
		}
	}
	for ; '0' <= b && b <= '9'; b = rc() {
		x = x*10 + int(b&15)
	}
	if neg {
		return -x
	}
	return
}

func (myio *MyIO) ReadInt64() (x int64) {
	rc := myio.rc
	b := rc()
	neg := false

	for ; '0' > b || b > '9'; b = rc() {
		if b == eof {
			return
		}
		if b == '-' {
			neg = true
		}
	}
	for ; '0' <= b && b <= '9'; b = rc() {
		x = x*10 + int64(b&15)
	}
	if neg {
		return -x
	}
	return x
}

func (myio *MyIO) ReadBytes() (s []byte) {
	rc := myio.rc
	b := rc()
	for ; b <= ' '; b = rc() {
	}
	for ; b > ' '; b = rc() {
		s = append(s, b)
	}
	return
}

func (myio *MyIO) ReadStr() (s string) {
	return string(myio.ReadBytes())
}

func (myio *MyIO) Flush() {
	myio.out.Flush()
}

func (myio *MyIO) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(myio.out, format, a...)
}

func (myio *MyIO) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(myio.out, a...)
}

func (myio *MyIO) Print(a ...any) (n int, err error) {
	return fmt.Fprint(myio.out, a...)
}

var myio *MyIO

func ri() int {
	return myio.ReadInt()
}

func ri64() int64 {
	return myio.ReadInt64()
}

func ria(_n int) []int {
	a := make([]int, _n)
	for i := range a {
		a[i] = ri()
	}
	return a
}

/* End of MyIO */

func sol() {

}

func main() {
	debug.SetGCPercent(-1)
	in, out := os.Stdin, os.Stdout
	myio = NewIO(in, out)
	defer myio.Flush()

	var t = 1
	t = ri()
	for ; t > 0; t-- {
		sol()
	}
}
