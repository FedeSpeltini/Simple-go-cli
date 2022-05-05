package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FedeSpeltini/go-cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("Enter expense: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	input = strings.Replace(input, "\n", "", -1)
	return input, nil
}

func ShowInConsole(expensesList []float32) {

	fmt.Println(contentString(expensesList...))
}

func Export(fileName string, expensesList []float32) error {
	f, e := os.Create(fileName)
	if e != nil {
		return e
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err := w.WriteString(contentString(expensesList...))
	if err != nil {
		return err
	}
	return w.Flush()
}

func expenseDetail(expensesList []float32) (max, min, sum, avg float32) {
	if len(expensesList) == 0 {
		return
	}
	max = expenses.Max(expensesList...)
	min = expenses.Min(expensesList...)
	sum = expenses.Sum(expensesList...)
	avg = expenses.Average(expensesList...)
	return
}

func contentString(expensesList ...float32) string {

	builder := strings.Builder{}
	max, min, sum, avg := expenseDetail(expensesList)
	fmt.Println("")
	for i, v := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense= %.2f\n", v))

		if i == len(expensesList)-1 {
			println("========================================================")
		}

	}
	builder.WriteString(fmt.Sprintf("Total= %.2f\n", sum))
	builder.WriteString(fmt.Sprintf("Max= %.2f\n", max))
	builder.WriteString(fmt.Sprintf("Min= %.2f\n", min))
	builder.WriteString(fmt.Sprintf("Average= %.2f\n", avg))
	return builder.String()
}
func ShowMin(expensesList []float32) {
	fmt.Println(expenses.Min(expensesList...))
}
