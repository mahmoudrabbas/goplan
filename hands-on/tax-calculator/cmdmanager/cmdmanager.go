package cmdmanager

import "fmt"

type CMDManger struct {
}

func New() CMDManger {
	return CMDManger{}
}

func (cmd CMDManger) ReadLines() ([]string, error) {
	fmt.Println("Please Enter prices, confirm every price with ENTER.")

	prices := []string{}

	for {
		var price string
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManger) WriteResults(data any) error {
	fmt.Println(data)
	return nil
}
