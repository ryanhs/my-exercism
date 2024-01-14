package greeting
import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	var out = "Hello, World!";
	fmt.Printf(out)
	return out
}
