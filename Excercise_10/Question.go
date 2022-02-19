package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var dayNumber int
	var day string
	var clientAge int
	var priceToPay float64
	var discountPercentage int
	const basePrice int = 9
	fmt.Println("*** Welcome to cinema ticket application ***")

	for {
		fmt.Println("Which day of the week are you going to cinema? (e.g enter 1 for Monday, 2 for Tuesday and ... 7 for Sunday) (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if day, error := strconv.Atoi(input); error == nil {
				if day >= 1 && day <= 7 {
					dayNumber = day
					break
				} else {
					fmt.Println("Error In The Entered Number! Try Again...")
				}
			} else {
				fmt.Println("Invalid Entry! Try Again...")
			}
		}
	}

	for {
		fmt.Println("How old are you? (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if age, error := strconv.Atoi(input); error == nil {
				if age >= 1 {
					clientAge = age
					break
				} else {
					fmt.Println("Error In The Entered Number! Try Again...")
				}
			} else {
				fmt.Println("Invalid Entry! Try Again...")
			}

		}
	}

	day, discountPercentage = getDayStringAndDiscountPercentage(dayNumber, clientAge)
	discountAmount := (discountPercentage * basePrice) / 100
	priceToPay = float64(basePrice - discountAmount)

	fmt.Println("Based on your Age(", clientAge, ") and the day you are going to the Cinema(", day, ") you have to pay $", priceToPay, " which is %", discountPercentage, "on discount($", discountAmount, ") in compare to base price ticket of $", basePrice)

}

func getDayStringAndDiscountPercentage(dayNumber int, clientAge int) (string, int) {
	var day string
	var discountPercentage int

	switch dayNumber {
	case 1:
		day = "Monday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 20
		} else {
			discountPercentage = 10
		}
		break
	case 2:
		day = "Tuesday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 0
		} else {
			discountPercentage = 30
		}
		break
	case 3:
		day = "Wednesday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 0
		} else {
			discountPercentage = 30
		}
		break
	case 4:
		day = "Thursday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 20
		} else {
			discountPercentage = 30
		}
		break
	case 5:
		day = "Friday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 0
		} else {
			discountPercentage = 10
		}
		break
	case 6:
		day = "Saturday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 0
		} else {
			discountPercentage = 10
		}
		break
	case 7:
		day = "Sunday"
		if clientAge >= 16 && clientAge <= 65 {
			discountPercentage = 0
		} else {
			discountPercentage = 10
		}
		break
	}
	return day, discountPercentage
}
