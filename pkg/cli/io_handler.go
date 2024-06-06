package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func HandleError(err error) {

}

type FileWriter struct {
	file *os.File
}

// Create New file writer
func NewFileWriter(file *os.File) *FileWriter {
	return &FileWriter{
		file: file,
	}
}

// Write implements the io.Writer interface
func (fw *FileWriter) Write(p []byte) (n int, err error) {
	return fw.file.Write(p)
}

func SetupConfig(logger *logrus.Logger, configName string) (configDir string) {
	// Check config in home
	homeDir, _ := os.UserHomeDir()
	configDir = path.Join(homeDir, ".terraforge")
	configFilePath := path.Join(configDir, "/", configName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Initialize default config
		logger.Info("Creating new config at ", configFilePath)
		err := InitializeConfig(logger, configDir, configFilePath)
		if err != nil {
			logger.Error(err)
		}
	} else if err != nil {
		logger.Panic(err)
	}
	return
}

func CreateAppStructure(appName string, dirs []string) (bool, error) {
	for _, dir := range dirs {
		err := os.MkdirAll(appName+dir, 0755)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// Check if Directory is Empty
// Returns (Empty Directory, Error)
func IsEmptyDir(name string) (os.FileInfo, error) {
	var (
		directory fs.FileInfo
		err       error
	)

	if directory, err = os.Stat(name); os.IsNotExist(err) {
		_ = os.Mkdir(name, 0755) // create new directory
		directory, _ = os.Stat(name)
	}

	// Check if it is a directory
	if !directory.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", name)
	}

	// Read the directory contents
	contents, err := os.ReadDir(directory.Name())
	if err != nil {
		return nil, fmt.Errorf("%s is not empty", name)
	}

	// Check if the directory is empty
	if len(contents) == 0 {
		return directory, nil
	}
	return nil, nil
}
