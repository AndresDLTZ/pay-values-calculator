package file_facade

import (
	"fmt"
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
