package main

import (
	"fmt"
	"github.com/andresdltz/pay-values-calculator/src/file_facade"
)

func main() {
	var filePath string
	fmt.Print("Ingrese ruta del archivo con nombre y horarios (.txt): ")
	// e.g files/name-and-schedules1.txt, files/name-and-schedules2.txt
	fmt.Scan(&filePath)

	// Define a file facade to then read and validate it, in order to calculate the values to be paid.
	f, err := file_facade.NewFileFacade(&filePath)
	if !(err != nil) {
		f.ReadFileAndCalculatePayValues()
	}

	defer func() {
		fmt.Println("\nPrograma terminado :)")
	}()
}
