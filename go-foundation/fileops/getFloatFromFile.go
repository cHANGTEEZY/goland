package fileops

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		fmt.Println("Errro!!!")
		return 1000.0, nil
	}
	balanceText := strings.TrimSpace(string(data))
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
