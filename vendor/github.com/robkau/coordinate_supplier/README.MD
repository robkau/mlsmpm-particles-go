# coordinate_supplier

[![Go Reference](https://pkg.go.dev/badge/github.com/robkau/coordinate_supplier.svg)](https://pkg.go.dev/github.com/robkau/coordinate_supplier)

coordinate_supplier is a utility Go package to supply coordinates in a XY grid.

----
## Features:
 - Hand out each coordinate exactly once, or loop through all coordinates repeatedly
 - Hand out coordinates in ascending order, descending order, or random order
 - Fast and mostly concurrent-safe (no data races) implementation via atomic.AddUint64 
 - Strictly concurrent-safe (guaranteed in order) implementation via sync.RWMutex
----

## Import and use:
```
import "github.com/robkau/coordinate_supplier"

// make supplier from options
opts := coordinate_supplier.CoordinateSupplierOptions{Width: 10, Height: 10, Order: coordinate_supplier.Asc, Repeat: false}
cs, err := coordinate_supplier.NewCoordinateSupplier(opts)
if err != nil {
  // handle me
}

// consume coordinates
for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
    fmt.Println("The next coordinate is", x, y)
}
```
----

## Run tests and benchmarks:
```
go test -bench=. -benchmem ./...
```
----

## Notes

Atomic implementation is typically 1.5x to 4x faster than RwMutex implementation:
```
goos: linux
goarch: amd64
pkg: github.com/robkau/coordinate_supplier
cpu: AMD Ryzen 5 1600 Six-Core Processor   

# 30x30 coordinates, 1 concurrent consumer, all items consumed once                 # bench iterations       # bench time          # bench allocs
BenchmarkCoordinateSuppliers/atomic-30w-30h-1consumers-consume1-12                	   29184	     42326 ns/op	   16560 B/op	       8 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-1consumers-consume1-12                    	   15271	     77926 ns/op	   16576 B/op	       8 allocs/op

# 30x30 coordinates, 10 concurrent consumers, all items consumed once
BenchmarkCoordinateSuppliers/atomic-30w-30h-10consumers-consume1-12               	   22893	     54415 ns/op	   17136 B/op	      17 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-10consumers-consume1-12                   	    4530	    271892 ns/op	   17162 B/op	      17 allocs/op

# 30x30 coordinates, 100 concurrent consumers, all items consumed once
BenchmarkCoordinateSuppliers/atomic-30w-30h-100consumers-consume1-12              	   10000	    100987 ns/op	   22896 B/op	     107 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-100consumers-consume1-12                  	    3457	    304284 ns/op	   22990 B/op	     107 allocs/op

# 30x30 coordinates, 1000 concurrent consumers, all items consumed once
BenchmarkCoordinateSuppliers/atomic-30w-30h-1000consumers-consume1-12             	    1803	    739194 ns/op	   80496 B/op	    1007 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-1000consumers-consume1-12                 	    1516	    770383 ns/op	   80851 B/op	    1010 allocs/op

# 30x30 coordinates, 1 concurrent consumer, on repeat until 1000000 items consumed
BenchmarkCoordinateSuppliers/atomic-30w-30h-1consumers-consume1000000-12          	      50	  21859851 ns/op	   16560 B/op	       8 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-1consumers-consume1000000-12              	      19	  53503156 ns/op	   16576 B/op	       8 allocs/op

# 30x30 coordinates, 10 concurrent consumers, on repeat until 1000000 items consumed
BenchmarkCoordinateSuppliers/atomic-30w-30h-10consumers-consume1000000-12         	      14	  75009663 ns/op	   17136 B/op	      17 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-10consumers-consume1000000-12             	       4	 310343362 ns/op	   17152 B/op	      17 allocs/op

# 30x30 coordinates, 100 concurrent consumers, on repeat until 1000000 items consumed
BenchmarkCoordinateSuppliers/atomic-30w-30h-100consumers-consume1000000-12        	      15	  69952752 ns/op	   22896 B/op	     107 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-100consumers-consume1000000-12            	       4	 259155838 ns/op	   24112 B/op	     119 allocs/op

# 30x30 coordinates, 1000 concurrent consumers, on repeat until 1000000 items consumed
BenchmarkCoordinateSuppliers/atomic-30w-30h-1000consumers-consume1000000-12       	      15	  71343118 ns/op	   80496 B/op	    1007 allocs/op
BenchmarkCoordinateSuppliers/rw-30w-30h-1000consumers-consume1000000-12           	       4	 287350628 ns/op	  102040 B/op	    1231 allocs/op
```
