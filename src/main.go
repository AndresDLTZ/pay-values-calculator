package main

import (
	"fmt"
	"github.com/andresdltz/pay-values-calculator/src/file_facade"
)

func main() {
	filePath := "files/name-and-schedules2.txt"
	// Define a file facade to then read and validate it, in order to calculate the values to be paid.
	f, err := file_facade.NewFileFacade(&filePath)
	if !(err != nil) {
		f.ReadFileAndCalculatePayValues()
	}

	defer func() {
		fmt.Println("\nPrograma terminado :)")
	}()
}
