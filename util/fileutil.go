package util

import (
	"io/ioutil"
	"sync"
)

/**
 *	Menulis ke File dan Baca Konten File
 *	Yang Aman dalam Multi-Goroutine (Multi-Threading)
 *
 *	@author Mufid Jamaluddin
 **/

// FileHandler : Handler IO File with Thread-Safe
type FileHandler struct {
	Filename string
}

// WriteToFile : Write to File Thread-Safe
func (fi *FileHandler) WriteToFile(content string) error {

	var mu sync.Mutex

	bytecontent := []byte(content)

	// Amanin Thread WriteFile, Lock dan Unlock
	mu.Lock()
	err := ioutil.WriteFile(fi.Filename, bytecontent, 0644)
	mu.Unlock()

	return err
}

// ReadFromFile : Read File Thread-Safe
func (fi *FileHandler) ReadFromFile() (string, error) {

	var mu sync.Mutex

	// Amanin Thread
	mu.Lock()
	bytecontent, err := ioutil.ReadFile(fi.Filename)
	mu.Unlock()

	content := string(bytecontent)

	return content, err
}
