package hsproject

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Project struct {
	filePath    string
	ProjectName string
	Description string
	Author      string
}

func CreateNew(fileName string) (*Project, error) {
	result := &Project{
		filePath:    fileName,
		ProjectName: "Untitled Project",
	}

	if err := result.Save(); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Project) Save() error {
	var err error
	var file []byte

	if file, err = json.MarshalIndent(p, "", "   "); err != nil {
		return err
	}
	if err = ioutil.WriteFile(p.filePath, file, os.FileMode(0644)); err != nil {
		return err
	}
	return nil
}

func LoadFromFile(fileName string) (*Project, error) {
	var err error
	var file []byte
	var result *Project

	if file, err = ioutil.ReadFile(fileName); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(file, &result); err != nil {
		return nil, err
	}

	result.filePath = fileName

	return result, nil
}