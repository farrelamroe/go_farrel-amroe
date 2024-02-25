package praktikum

import "fmt"

func sum(n int) int {
	result := 0
	for i := n; i > 0; i-- {
		for j := 0; j < n; j++ {
			result += i;
		}
	}
	return result
}
func main() {

}
