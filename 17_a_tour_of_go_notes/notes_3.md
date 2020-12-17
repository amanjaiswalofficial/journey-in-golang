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