package alias

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Alias struct {
	FilePath string
}

// NewAlias creates a new Alias struct to handle all
func NewAlias(path string) *Alias {
	return &Alias{FilePath: path}
}

func (a *Alias) UpdateAliasFile(name string, command string) error {
	file, err := os.OpenFile(a.FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}

	exists, err := a.CheckIfAliasExistAndUpdate(name, command, file)

	if err != nil {
		log.Fatalf("Error Found: %s", err)
	}

	if !exists {
		_, err = file.WriteString("alias " + name + "='" + command + "'")
	}

	return err
}

func (a *Alias) CheckIfAliasExistAndUpdate(name string, command string, file *os.File) (bool, error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "alias "+name) {
			a.ReplaceAlias(name, command, file.Name())
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func (a *Alias) ReplaceAlias(name, command string, fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Error while loading File Content: %s", err)
	}

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, name) {
			lines[i] = "alias " + name + "='" + command + "'"
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalf("Error while write to File: %s", err)
	}
}
