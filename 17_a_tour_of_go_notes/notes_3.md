### Methods on non-struct types (int, float etc)

```
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

Here MyFloat is a struct over type float64, so a reciever method over MyFloat can be used and indirectly over float value. Similar can be used for int str etc type of values

### Pointer Receivers

```
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

As pointers allow to access and modify the original value instead of a copy of original, used more than normal variable receivers.

### Pointer Indirectionn in functions

```
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	(&v).Scale(3) // This works as per pointer receiver
	v.Scale(2) // But, this also works, as Go automatically detects it

	fmt.Println(v)
}

```

Here, even though the method Scale is a receiver on a pointer of Vertex, `v` and `&v` both are accepted as valid pointer receiver for Scale() method.

### Choosing pointer receivers

Reasons to use pointer reciever over value receiver

1.The method modifies the actual value and not its copy.
2.Avoid copying value on each method call, which can be expensive on large structs.

### Interface In Go
Set of method signatures, as per A Tour of Go

```
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	a = v // *v implements Abser, v doesn't hence not interface

	fmt.Println(a.Abs()) // won't work since pointing to v
}
```

Here, `MyFloat` and `Vertex` both implement `Abs()`, as per required by the definition of Abser Interface, hence both `f` & `*v` are interfaces whereas no such method that accepts `v`as input, hence v isn't an interface.


### Interface Values

```
func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %
	T)\n", i, i)
}
```
Beneath the definition, an interface is more or less a tuple of (Value, Type)

Returning output as

```
(&{Hello}, *main.T)
Hello
```

### Interface for nil underlying values

If the value inside the interface is nil, then the interface will be called over a nil value
```
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	// will be called on a nil receiver
	i = t
	i.M()

	i = &T{"hello"}
	i.M()
}
```

### Using godoc for help on packages

```
godoc -http :6000
```

### Using go build to make binaries

Inside the folder containing .go file
```
go build
```

To generate a build for Windows

```
GOOS=windows go build
```

To see documentation on any of the packages/methods

```
godoc -http 6060
```

