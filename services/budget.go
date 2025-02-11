package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SetBudgetAction struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func BudgetMain() {
	fmt.Print(Divider)
	fmt.Println("You are in the \"Budget\" section. Choose one of the following options:\n")

	CyanPrint("- \"%s\" to create a new monthly budget\n", "create")
	CyanPrint("- \"%s\" to view your current monthly budget\n", "view")
	CyanPrint("\nType \"%s\" to return to the main menu.\n", "back")
	fmt.Print(Divider)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())
		switch cmd {
		case "create":
			CreateBudgetPage()
		case "view":
			ViewBudget()
		case "back":
			WelcomeText()
			return
		}
	}
}

func CreateBudgetPage() {
	fmt.Print(Divider)
	fmt.Println("You are in the \"Budget\" section. Choose one of the following:")
	fmt.Println("\n- Enter your total monthly budget: (e.g., 2000)")
	CyanPrint("After setting the amount, you can \"%s\" the last operation\n", "undo")
	fmt.Print(Divider)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		budgetStr := strings.TrimSpace(scanner.Text())
		budget, err := strconv.Atoi(budgetStr)
		createBudget(budget)

		if err != nil {
			fmt.Println("Please enter a valid budget amount")
			continue
		}

		CyanPrint("\nYour monthly budget of %s has been successfully set!\n", budgetStr)

		CyanPrint("\nType \"%s\" cancel the setting of this monthly budget or ", "undo")
		CyanPrint("\"%s\" to return to the main menu.\n", "back")
		fmt.Print(Divider)

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

func createBudget(budget int) {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strings.ToLower(now.Month().String())
	//var oldBudget int
	//var newBudget int
	var newPeriod Month

	data := LoadData()

	if _, ok := data.Years[year]; !ok {
		newPeriod = data.Years[year].Months[month]

		//oldBudget = 0
		//newBudget = budget

		newPeriod.Budget = budget
		newPeriod.Expenses = make([]Expense, 0)

		fmt.Println(newPeriod)
	} else {
		newPeriod = data.Years[year].Months[month]

		//oldBudget = newPeriod.Budget
		//newBudget = budget

		newPeriod.Budget = budget
	}

	//data.LastAction = Action{Type: "set_budget", Data: SetBudgetData{oldBudget, newBudget, year, month}}

	data.Years[year].Months[month] = newPeriod

	err := SaveData(data)

	if err != nil {
		fmt.Println(err)
	}
}

func ViewBudget() {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strings.ToLower(now.Month().String())

	data := LoadData()
	currentMonthData := data.Years[year].Months[month]
	budget := currentMonthData.Budget
	var remainingBudget int
	var expenses int

	for _, exp := range currentMonthData.Expenses {
		expenses += exp.Amount
	}

	remainingBudget = budget - expenses

	fmt.Print(Divider)
	fmt.Println("View your current monthly budget\n")
	fmt.Printf("Your current monthly (%s %v) budget is:\n", month, year)
	fmt.Printf("- Total Budget: %v\n", budget)
	fmt.Printf("- Remaining Budget: %v\n\n", remainingBudget)

	fmt.Println("To view another month, type:")
	CyanPrint("- Month and year (%s)\n", "e.g., \"september 2024\"")

	CyanPrint("\nOr type \"%s\" to return to the main menu.\n", "back")
	fmt.Print(Divider)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()
	cmd := strings.TrimSpace(scanner.Text())

	if cmd == "back" {
		MainScreen()
		return
	} else {
		month, year, ok := ExtractMonthYear(cmd)

		if ok {
			monthData := data.Years[year].Months[month]

			budget := monthData.Budget
			var remainingBudget int
			var expenses int

			for _, exp := range monthData.Expenses {
				expenses += exp.Amount
			}

			remainingBudget = budget - expenses

			fmt.Printf("- Total Budget: %v\n", budget)
			fmt.Printf("- Remaining Budget: %v\n\n", remainingBudget)

			fmt.Print("> ")
			scanner.Scan()
			cmd := strings.TrimSpace(scanner.Text())

			if cmd == "back" {
				MainScreen()
			}
		}
	}

}
