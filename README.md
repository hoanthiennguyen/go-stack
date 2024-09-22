# Stack
Since Go's standard package does not provide stack out of the box, I've decided to create an interface and an implementation for it so that I don't need to create a stack from scratch every time I need it.
I hope that some of you find it useful as well.

## Usage
```
import (
	"fmt"

	"github.com/hoanthiennguyen/go-stack"
)

func main() {
	stk := stack.New[int]()
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Pop()
	last, _ := stk.Peek()
	fmt.Printf("top: %d\n", last)
}
```
