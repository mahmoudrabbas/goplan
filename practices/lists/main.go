package lists

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// 1
	hoppies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hoppies)
	println("__________________________")
	// 2
	fmt.Println(hoppies[0])
	fmt.Println(hoppies[1:3])
	println("__________________________")
	// 3
	mainHoppies := hoppies[1:]
	fmt.Println(mainHoppies)
	println("__________________________")
	// 4
	fmt.Println("Cap: ", cap(mainHoppies))
	mainHoppies = mainHoppies[0:]
	fmt.Println(mainHoppies)
	println("__________________________")

	// 5
	courseGoals := []string{"Learn Go", "Practice it"}
	fmt.Println(courseGoals)
	println("__________________________")
	// 6
	courseGoals[0] = "Rest Api"
	println(cap(courseGoals))
	courseGoals = append(courseGoals, "GoLang")
	fmt.Println(courseGoals)
	println("__________________________")

	// 7

	products := []Product{{id: "1", title: "p1", price: 50.0}, {id: "2", title: "p2", price: 100.00}}
	fmt.Println(products[0])

	products2 := []Product{
		{id: "3", title: "p3", price: 555.5},
		{id: "4", title: "p4", price: 555.5},
	}

	products = append(products, products2...)
	fmt.Println(products)

}
