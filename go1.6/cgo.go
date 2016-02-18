package main

/*
int fn(void* arg) { return arg == 0; }
*/
import "C"
import "unsafe"

type T struct{ a, b int }
type X struct{ t *T }

func main() {
	t := T{a: 1, b: 2}
	x := X{t: &t}

	C.fn(unsafe.Pointer(&t))
	C.fn(unsafe.Pointer(&x))
}
