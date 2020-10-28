package main

import "gosl/io"

func main() {
	io.Pf("Hello Gosl (6031)!!!\n")
	for i, n := range []int{6, 0, 3, 1} {
		io.Pf("%d: the number is now = %d\n", i, n)
	}
}
