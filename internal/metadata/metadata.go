package metadata

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type trackMetadata struct {
	song   string
	artist string
	album  string
	year   string
}

// GetMetadata returns matching metadata in order of most relevant
func GetMetadata(filename string) ([]string, error) {
	fmt.Println("Path to file", filename)
	validFile, invalidFileReason := validFile(filename)
	if !validFile {
		return nil, errors.New(invalidFileReason)
	}
	return []string{"output"}, nil
}

func validFile(filename string) (bool, string) {
	var validFileExtensions = []string{".mp3", ".m4a"}
	valid := true
	var invalidFileReason string
	if !fileExists(filename) {
		valid = false
		invalidFileReason = "File does not exist"
	} else {
		extension := filepath.Ext(filename)
		if (!stringInSlice(extension, validFileExtensions)) {
			valid = false
			invalidFileReason = "Not a valid file extension - only " + strings.Join(validFileExtensions, ", ") + " supported."
		}
	}
	return valid, invalidFileReason
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return false
	}
	return true
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
