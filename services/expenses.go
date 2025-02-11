package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExpensesMain() {
	fmt.Print(Divider)
	fmt.Println("You are in the \"Add Expenses\" section.\n")
	fmt.Println("Please enter your expense category and amount. Added expense will be deducted from budget.")
	fmt.Println("After entering the expense amount, you will have the option to undo it with the \"undo\" command.\n")

	scanner := bufio.NewScanner(os.Stdin)

	// Get the category from the user
	fmt.Println("Category (e.g., Food):")
	fmt.Print("> ")
	scanner.Scan()
	category := strings.TrimSpace(scanner.Text())

	// Get the amount from the user
	fmt.Println("Amount (e.g., 100):")
	fmt.Print("> ")
	scanner.Scan()
	amount, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if err != nil {
		fmt.Println("Invalid amount.")
	}

	// Add expense to DB
	AddExpense(category, amount)

	// Show the result to the user
	fmt.Printf("\nYour expense of %v in category %s has been successfully added!\n", amount, category)
	CyanPrint("\nType \"%s\" cancel the setting of this monthly budget or ", "undo")
	CyanPrint("\"%s\" to return to the main menu.\n", "back")
	fmt.Print(Divider)

	for {
		fmt.Print("> ")
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())

		switch cmd {
		case "undo":
			Undo()
		case "back":
			MainScreen()
			return
		}
	}
}

func AddExpense(category string, amount int) {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strings.ToLower(now.Month().String())
	var period Month

	data := LoadData()

	if _, ok := data.Years[year]; !ok {
		period = data.Years[year].Months[month]
		period.Budget = 0
		period.Expenses = make([]Expense, 0)
	} else {
		period = data.Years[year].Months[month]
	}

	period.Expenses = append(period.Expenses, Expense{Category: category, Amount: amount})

	data.Years[year].Months[month] = period
	data.LastAction = Action{Type: "add_expense", Data: AddExpenseData{amount, category, year, month}}

	err := SaveData(data)

	if err != nil {
		fmt.Println("Error saving data:", err)
	}
}
