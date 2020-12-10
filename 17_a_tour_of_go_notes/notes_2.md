### Working with pointer
Similar to C, `&` is used to access the address of a variable
To access the same address via pointer, `*` is used
This is called **dereferencing**.

```
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

### Using Pointers with structs

Usually to access any variable via pointer, `*` is used
But in case of structs, `*` can be ommitted and would still work.

```
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	(*p).X = 10 // basic way to do so
	p.X = 20 // without having to use *
	fmt.Println(v)
}
```

### Slice Literals

Using a slice literal for defining arrays is like defining an array without length, which is then assumed if any elements are provided while defining.

So,

Ex - 
`[3]bool{true, false, false}` is same as `[]bool{true, false, false}`

### Using make for slices

```
a := make([]int, 5)

b := make([]int, 0, 5)
```

### Type Assertion
A type assertion provides access to an interface value's underlying concrete value. 

```
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

```
An empty interface can have any kind of value
When there's enough information about what kind of value it's storing

A type assertion (Ex- `i.(string)`) can be used to get the value out

Here, since it was known that the value inside the interface is string, a type assertion for string was used.