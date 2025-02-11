package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Action struct {
	Type string         `json:"type"`
	Data AddExpenseData `json:"data"`
}

type SetBudgetData struct {
	OldBudget int    `json:"old_budget"`
	NewBudget int    `json:"new_budget"`
	Year      string `json:"year"`
	Month     string `json:"month"`
}

type AddExpenseData struct {
	Amount   int    `json:"amount"`
	Category string `json:"category"`
	Year     string `json:"year"`
	Month    string `json:"month"`
}

type Expense struct {
	Amount   int    `json:"amount"`
	Category string `json:"category"`
}

type Month struct {
	Budget   int       `json:"budget"`
	Expenses []Expense `json:"expenses"`
}

type Year struct {
	Months map[string]Month `json:"months"`
}

type Data struct {
	LastAction Action          `json:"last_action"`
	Years      map[string]Year `json:"years"`
}

const filename = "data.json"

func SaveData(data Data) error {
	file, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}

func LoadData() Data {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	file, err := os.ReadFile(path.Join(dir, filename))

	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	var data Data
	err = json.Unmarshal(file, &data)

	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	return data
}
