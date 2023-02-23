import (
  "fmt"
)  

func remove(x int, y int) int {
	total := 0
	total = x - y
	return total
}

func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func printRes(int res) {
  fmt.Sprintf("hello test")
    fmt.Sprintf(%s,res)
}
func main() { 
  printRes(add(2,3))
}
