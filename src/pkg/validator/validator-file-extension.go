package validator

import (
	"errors"
	"os"
)

var errFileExtension = errors.New("extension del archivo debe ser .txt")

// FileExtensionValidator type provides a structure that takes care of getting file and file extension
type FileExtensionValidator struct {
	file *os.File
	fileExtension string
}

// NewFileExtensionValidator will retrieve a file and file extension and instance a FileExtensionValidator type
// file extensions, e.g txt, go, py
func NewFileExtensionValidator(file *os.File, fileExtension string) FileExtensionValidator {
	return FileExtensionValidator{file, fileExtension}
}

// ValidateFileExtension validates the file extension
// file extensions, e.g txt, go, py
func (vfe *FileExtensionValidator) ValidateFileExtension() error {
	r := getRegexpCompile("(.|\\w|)+\\." + vfe.fileExtension + "$")
	if r.MatchString(vfe.file.Name()) != true { return errFileExtension }
	return nil
}