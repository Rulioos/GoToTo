package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

//scanProject scans a whole project and returns all go filepaths in the path tree.
func scanProject(scanPath string) ([]string, error) {
	goFilesPath := make([]string, 0)
	e := filepath.Walk(scanPath, func(path string, f os.FileInfo, err error) error {
		if ext := strings.SplitN(path, ".", 2); len(ext) == 2 && ext[1] == "go" {
			goFilesPath = append(goFilesPath, path)
		}

		return err
	})
	if e != nil {
		return nil, e
	}

	return goFilesPath, nil
}

//YamlConf holds the configuration that will drive ts files creation.
type YamlConf struct {
	OutputDirPath       string
	ContextList         []string
	FilenameContextsMap map[string][]string
	Batches             map[string][]TsInterface
}

//GenerateConfigYaml generate a GototoConf.yaml that will drive the generateTS command.
func GenerateConfigYaml(scanPath string, outputDir string) {
	files, err := scanProject(scanPath)
	if err != nil {
		panic(err)
	}

	//Parsing
	contextSlice := make([]string, 0)
	batches := make(map[string][]TsInterface)
	for _, gof := range files {
		localBatch := ParseGoFile(gof, &contextSlice)

		for c, i := range localBatch {
			if _, keyInMap := batches[c]; keyInMap {
				batches[c] = append(batches[c], i...)
			} else {
				batches[c] = i
			}

		}
	}

	SetifyString(&contextSlice)

	//Building conf
	yamlStruct := YamlConf{OutputDirPath: outputDir,
		FilenameContextsMap: make(map[string][]string),
		ContextList:         contextSlice,
		Batches:             batches}
	for _, s := range contextSlice {
		l := make([]string, 0)
		l = append(l, strings.ToLower(s))
		yamlStruct.FilenameContextsMap[strings.Title(s)+"Model"] = l
	}

	//Marshalling contexts
	yamlStructConf, err := yaml.Marshal(yamlStruct)
	Check(err)
	//writing file
	fconf, err := os.Create(scanPath + "/gototoConf.yaml")
	Check(err)
	_, err = fconf.Write(yamlStructConf)
	Check(err)
	fconf.Close()
}

//GetYamlConfig gets the Yaml configuration from the file if it does exists.
func (y *YamlConf) GetYamlConfig(scanPath string) (*YamlConf, error) {
	configPath := scanPath + "/gototoConf.yaml"
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, y)
	Check(err)
	return y, nil
}
