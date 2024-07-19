package watcher

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShouldDetectChanges(t *testing.T) {
	folderToWatch := filepath.Join("..", "example")
	fileToCreate := "temp.js"
	fileToCreateFullPath := filepath.Join(folderToWatch, fileToCreate)

	w := New(folderToWatch)

	c := make(chan string)
	go w.Read(c)

	file, err := os.Create(fileToCreateFullPath)
	if err != nil {
		t.Errorf("file %s is not created", fileToCreateFullPath)
		return
	}

	content := "Hello, this is a sample text."
	_, err = file.WriteString(content)
	if err != nil {
		t.Errorf("file %s is not writted", fileToCreateFullPath)
		return
	}

	file.Close()

	fileName := <-c

	err = os.Remove(fileToCreateFullPath)
	if err != nil {
		t.Errorf("file %s is not deleted", fileToCreateFullPath)
		return
	}

	if fileName != fileToCreate {
		t.Errorf("detected change on %s instead of %s", fileName, fileToCreate)
	}
}
