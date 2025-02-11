package services

import (
	"fmt"
	"github.com/fatih/color"
	"regexp"
)

func CyanPrint(s, w string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(s, cyan(w))
}

func Undo() {
	data := LoadData()
	lastAction := data.LastAction

	if lastAction.Type == "" {
		fmt.Println("No history of actions")
	} else {
		switch lastAction.Type {
		case "add_expense":
			var expense = lastAction.Data
			monthlyBudget := data.Years[expense.Year].Months[expense.Month]
			monthlyBudget.Expenses = monthlyBudget.Expenses[:len(monthlyBudget.Expenses)-1]

			data.Years[expense.Year].Months[expense.Month] = monthlyBudget
		}
		data.LastAction = Action{}
		err := SaveData(data)

		if err != nil {
			fmt.Println("Error saving data:", err)
		}

		fmt.Println("\nThe last action was successfully undone.")
	}
}

func ExtractMonthYear(s string) (string, string, bool) {
	re := `^(january|february|march|april|may|june|july|august|september|october|november|december)\s(\d{4})$`
	rex := regexp.MustCompile(re)
	match := rex.FindStringSubmatch(s)

	if len(match) == 3 {
		return match[1], match[2], true
	}

	return "", "", false
}
