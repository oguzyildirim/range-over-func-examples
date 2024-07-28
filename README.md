# Range over func

### Contents

* [What Changed in golang 1.21](#What Changed in golang 1.21)
* [Range Over Functions](#Range Over Functions)
* [Examples](#examples)

### What Changed in golang 1.21

- Different loop variable each iteration (GOEXPERIMENT=loopvar)
```go
    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
        }()
    }
```

- Range over integers
```go
    for i := range 23 {
		fmt.Println(i)
	}
```

### Range Over Functions

+ Go 1.22 introduces an experimental feature: range over functions. This new construct allows you to iterate over values produced by a function, similar to how you can iterate over elements of a slice or map. This feature brings more flexibility and expressiveness to Go's iteration capabilities.

**Motivation**

+ since generics the data structure decoupled the types (allows custom containers)
+ no standard way to iterate over a sequence of values in Go
+ easy to move around in a large code base

#### Example
```go
func Backward[E any](s []E) func(func(int, E) bool) {
    return func(yield func(int, E) bool) {
        for i := len(s)-1; i >= 0; i-- {
            if !yield(i, s[i]) {
            return
            }
        }
    }
}
```


**Mechanics**

+ Signatures
```go
type Seq[V any] func(yield func(V) bool)
type Seq2[K, V any] func(yield func(K, V) bool)

func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func())
func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())
```


+ Compiler Transformation
```go
package slices

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}
```

```go
s := []string{"hello", "world"}
for i, x := range slices.Backward(s) {
    fmt.Println(i, x)
}
```

```go
slices.Backward(s)(func(i int, x string) bool {
    fmt.Println(i, x)
    return true
})
```


**Performance**
+ one item at a time
+ no allocation
+ early return


### Examples

#### Generator

```go
func ToInfinity(start, step int) iter.Seq[int] {
	if step > 0 {
		// Forward iteration from start
		return func(yield func(int) bool) {
			for i := start; ; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}

	if step < 0 {
		// Backward iteration from start
		return func(yield func(int) bool) {
			for i := start; ; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}

	return func(yield func(int) bool) {
		// yield only the 'start' value when step 0
		yield(start)
	}
}
```

#### Timer

```go
func every(d time.Duration) func(func() bool) {
	return func(yield func() bool) {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if !yield() {
					return
				}
			}
		}
	}
}
```

#### File Line Reader

```go
func LineReader(filename string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if !yield(line) {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning file: %s\n", err)
			return
		}
	}
}
```

#### Loop

+ https://github.com/dreamsofcode-io/loop

```go
    xs := []int{1,2,3,4,5}
    for _, x := range loop.Pool(xs, 2) {
        time.Sleep(time.Second)
    }
	
	
    nums := []int{1, 2, 3, 4, 5}
    for i, batch := range loop.Batch(nums, 2) {
        fmt.Println(i, batch)
    }
```

#### Sqlrange

+ https://github.com/achille-roussel/sqlrange

```go
type Point struct {
    X float64 `sql:"x"`
    Y float64 `sql:"y"`
}
for p, err := range sqlrange.Query[Point](db, `select x, y from points`) {
    if err != nil {
        ...
    }
    ...
}
```

#### Kway-go

+ https://github.com/achille-roussel/kway-go

```go
for v, err := range kway.Merge(seq0, seq1, seq2) {
    fmt.println(value)
}
```


### More

#### Proposal
- https://github.com/golang/go/discussions/56413

#### Issues
- https://github.com/golang/go/issues/61405
- https://github.com/golang/go/issues/61898
- https://github.com/golang/go/issues/61899
- https://github.com/golang/go/issues/61900
- https://github.com/golang/go/issues/61901
- https://github.com/golang/go/issues/61902

#### Wiki
- https://go.dev/wiki/RangefuncExperiment

#### New range func
- https://tip.golang.org/src/cmd/compile/internal/rangefunc/rewrite.go
