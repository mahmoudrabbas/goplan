package maps

import "fmt"

type floatMap map[int]string

func (f floatMap) get() {
	fmt.Println(f)
}

func main() {

	// var fm floatMap = make(floatMap)
	// fm[1] = "DD"
	// fm[2] = "dd"
	// fm[3] = "ss"
	// fm.get()

	// for k, v := range fm {
	// 	fmt.Println("Key", k, "Value", v)
	// }

	slice := make([]int, 2, 5)
	slice[0] = 5
	slice[1] = 5
	fmt.Println(slice)
}
