package main

import (
	"fmt"

	"example.com/first-app/user"
)

func getUserData(prompt string) (val string) {
	fmt.Print(prompt)

	fmt.Scanln(&val)
	return
}

func main() {

	// firstName := getUserData("Please Enter you firstName: ")
	// lastName := getUserData("Please Enter you lastName: ")
	// date := getUserData("Please Enter you date: ")

	// person, err := user.New(firstName, lastName, date)

	admin := user.NewAdmin("mahmoud", "123")

	admin.OutputUserData()
	admin.ClearUserName()
	admin.OutputUserData()

	// if err != nil {
	// 	fmt.Print(err)
	// } else {
	// 	person.OutputUserData()

	// 	person.ClearUserName()

	// 	person.OutputUserData()
	// }

}
