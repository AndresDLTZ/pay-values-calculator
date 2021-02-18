package file_facade

import (
	"bufio"
	"fmt"
	"github.com/andresdltz/pay-values-calculator/src/pkg/line_name_schedule"
	"github.com/andresdltz/pay-values-calculator/src/pkg/pay_value"
	"github.com/andresdltz/pay-values-calculator/src/pkg/validator"
	"os"
)

// FileFacade type provides a structure that takes care of getting file path,
// validating the extension and getting the file
type FileFacade struct {
	filePath               string
	file                   *os.File
	fileExtensionValidator validator.FileExtensionValidator
}

// NewFileFacade will retrieve a file path and instance a FileFacade type
func NewFileFacade(filePath *string) (FileFacade, error) {
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error: %s no existe tal archivo o directorio\n", *filePath)
		return FileFacade{}, err
	}
	return FileFacade{
		filePath:               *filePath,
		file:                   file,
		fileExtensionValidator: validator.NewFileExtensionValidator(file, "txt"),
	}, nil
}

// ReadFileAndCalculatePayValues reads the file line by line and then calculates the values to be paid.
func (f *FileFacade) ReadFileAndCalculatePayValues() {
	if err := f.fileExtensionValidator.ValidateFileExtension(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	scanner := bufio.NewScanner(f.file)

	// lineCh channel receives iterated lines from the file to send them to goroutine CalculatePayValue
	lineCh := make(chan *line_name_schedule.Line)
	// doneCh channel sends a message when a goroutine has been completed
	doneCh := make(chan *bool)

	fmt.Println()
	go pay_value.CalculatePayValue(lineCh, doneCh)

	for scanner.Scan() {
		lineStr := scanner.Text()
		if err := validator.ValidateNameSchedule(&lineStr); err != nil {
			fmt.Printf("%s\nAdvertencia: %s\n\n", lineStr, err)
			continue
		}
		line := line_name_schedule.SplitNameAndSchedule(&lineStr)
		lineCh <- &line
	}
	close(lineCh)
	<- doneCh
}
