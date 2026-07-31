# Advanced Go Math


Is a initial beginner project that implements simple mathematicals formulas in `Go`.

## Prerequisites

> Go 1.18 or newer

## Installation
```bash
go get github.com/Minhagatasodorme/Advanced-Go-Math
```

## Example for use
```go
package main

import (
	"fmt"

	gm "github.com/Minhagatasodorme/Advanced-Go-Math/gomath"
)

func main() {
	// gm.Bhaskara(A, B, C)
	x1, x2, err := gm.Bhaskara(1, -5, 6)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The roots of this equation is %.2f and %.2f.\n", x1, x2)
}
```

**Output:**

```text
The roots of this equation is 3.00 and 2.00.
```

## Note
This is an initial project created to learn more about Go, so it currently includes only one mathematical formula. More updates coming soon!
