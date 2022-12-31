package main

// https://go.dev/doc/faq#stack_or_heap 
//
// Go compilers will allocate variables that are local to a function in that
// function's stack frame. However, if the compiler cannot prove that the variable
// is not referenced after the function returns, then the compiler must allocate the
// variable on the garbage-collected heap to avoid dangling pointer errors. Also,
// if a local variable is very large, it might make more sense to store it on the heap
// rather than the stack.
//
// If a variable has its address taken, that variable is a candidate for allocation
// on the heap. However, a basic escape analysis recognizes some cases when such
// variables will not live past the return from the function and can reside on the
// stack.

type BigStruct struct {
    A, B, C int
    D, E, F string
    G, H, I bool
}

func main() {
    println("foo");
}

func main1() {
    _ = stackIt()
}

//go:noinline
func stackIt() int {
    y := 2
    return y * 2
}

func main2() {
    _ = stackIt2()
}

//go:noinline
// https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
//
// stackIt2 passes the address of `y` up the stack to `main2`, where `y` will
// be referenced after the stack frame of `stackIt2` has already been freed.
// The compiler is therefore able to judge that `y` must be moved to the heap
// to remain alive. If it doesn't do this, we'll get a nil pointer in `main2`
// when attempted to reference `y`.
func stackIt2() *int {
    y := 2
    res := y * 2
    return &res
}

func main3() {
    y := 2
    _ = stackIt3(&y) // pass y down the stack as a pointer
}

//go:noinline
// https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
//
// stackIt3 on the other hand, passes `y` down the stack, and `y` isn't referenced
// anywhere outside `main3`. The compiler is therefore able to judge that `y` can
// exist within the stack alone, and doesn't need to be allocated to the heap.
// We won't be able to produce a nil pointer in any circumstances by referencing
// `y`.
func stackIt3(y *int) int {
    res := *y * 2
    return res
}

func copyIt() {
    _ = CreateCopy()
}

func pointerIt() {
    _ = CreatePointer()
}

//go:noinline
// CreateCopy copies `BigStruct` instances between stack frames.
func CreateCopy() BigStruct {
    return BigStruct{
        A: 123, B: 456, C: 789,
        D: "ABC", E: "DEF", F: "HIJ",
        G: true, H: true, I: true,
    }
}

//go:noinline
// CreatePointer shares `BigStruct` pointers up the stack, avoiding copying
// but resulting in heap allocations.
func CreatePointer() *BigStruct {
    return &BigStruct{
        A: 123, B: 456, C: 789,
        D: "ABC", E: "DEF", F: "HIJ",
        G: true, H: true, I: true,
    }
}
