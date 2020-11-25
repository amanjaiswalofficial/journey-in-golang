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