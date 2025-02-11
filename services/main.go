package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Feature1    = "Create and view your monthly budget"
	Feature2    = "Add expenses to your budget"
	BudgetCmd   = "budget"
	ExpensesCmd = "expenses"
)

func MainScreen() {
	scanner := bufio.NewScanner(os.Stdin)
	WelcomeText()

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "back":
			WelcomeText()
		case "budget":
			BudgetMain()
		case "expenses":
			ExpensesMain()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Sorry, I don't know that command.")
		}
	}
}

func WelcomeText() {
	fmt.Print(Divider)
	fmt.Println("Welcome to Budget Planner!")
	fmt.Println("Easily manage your monthly budget with simple commands.")
	fmt.Println("\nIn this version, you can:")
	fmt.Println(fmt.Sprintf("- %s", Feature1))
	fmt.Println(fmt.Sprintf("- %s", Feature2))
	fmt.Println("\nType one of the following commands to navigate:")
	CyanPrint("- \"%s\" to manage your monthly budget (create/view)\n", BudgetCmd)
	CyanPrint("- \"%s\" to manage your expenses (add)\n", ExpensesCmd)
	fmt.Print(Divider)
}
