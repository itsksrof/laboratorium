# Allocations

Benchmarks that compare copying versus using pointers.

```text
BenchmarkStackIt-12      	851856218	        1.401 ns/op	      0 B/op	      0 allocs/op
BenchmarkStackIt2-12     	82259503	       13.96 ns/op	      8 B/op	      1 allocs/op
BenchmarkStackIt3-12     	822049893	        1.360 ns/op	      0 B/op	      0 allocs/op
BenchmarkCopyIt-12       	251736740	        4.601 ns/op	      0 B/op	      0 allocs/op
BenchmarkPointerIt-12    	27371060	       43.33 ns/op	     80 B/op	      1 allocs/op
```

# References
- [Understanding Allocations in Go](https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d)
