package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateDir(dir string) error {
	newpath := filepath.Join(".", dir)
	err := os.MkdirAll(newpath, os.ModePerm)
	return err
}

func RmDir(dir string) error {
	path := filepath.Join(".", dir)
	err := os.RemoveAll(path)
	return err
}

func StatDir(dir string) (os.FileInfo, error) {
	path := filepath.Join(".", dir)
	return os.Stat(path)
}

func Write(file string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(file, data, perm)
	return err
}

func Read(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
