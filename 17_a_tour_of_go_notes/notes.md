### Exported Names
Any import starting with capital letter, is a name that was exported, which is different from a usual import (*i.e. starting with small letter*).

###### Ex - Pi from math is an exported name

So, pi won't work in place of Pi
```
package main

import (
	"fmt"
	"math"
)

func main() {
	// Won't work
	//fmt.Println(math.pi)
	// Will work
	fmt.Println(math.Pi)
}

```

### Multiple same type args in functions

While defining a function, when two or more consecutive named function parameters share a type, you can omit the type from all but the last. 

Ex - `x int, y int` can be `x, y int`


### Named return values

While defining the method, the values named on top can simply be returned in order they were written without explicitly returning it in the end.


A return statement without arguments returns the named return values. This is known as a "**naked**" return. 

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return here returns 2 values, x and y in same order
	return
}

func main() {
	// wil print 7 10 as values x y are returned
	fmt.Println(split(17))
}

```

### Variables and their postions

`var` can be at function level or package level.

```
package main

import "fmt"

// package level
// type at the end, same for all variables
var c, python, java bool

func main() {
	// function level
	var i int
	fmt.Println(i, c, python, java)
}

```

### Variables with initializers

No need of type required, the the variable will take the type of the initializer.

`var c, python, java = true, false, "no!"`
