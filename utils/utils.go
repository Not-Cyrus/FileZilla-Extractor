package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile(src, dst string) error {
	fileExists, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fileExists.Mode().IsRegular() {
		return fmt.Errorf("%s is not a file", err)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(fmt.Sprintf("%s/%s", CurrentPath, dst))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)

	return err
}

func MakeDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.Mkdir(dirName, 0700)
	}
	return nil
}

func OpenFile(fileName string) (fileData []byte, err error) {
	fileData, err = ioutil.ReadFile(fileName)
	return
}

func WriteJSON(fileName string, data interface{}) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/results/%s", CurrentPath, fileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := new(bytes.Buffer)
	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")

	err = enc.Encode(data)
	if err != nil {
		return err
	}

	_, err = file.Write(w.Bytes())
	if err != nil {
		return err
	}

	return nil
}

var (
	CurrentPath, _ = os.Getwd()
)
