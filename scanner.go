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
	Pending             []string
	FilenameContextsMap map[string][]string
	BatchesInterface    map[string][]TsInterface
	BatchesEnums        [][]Enum
}

//GenerateConfigYaml generate a GototoConf.yaml that will drive the generateTS command.
func GenerateConfigYaml(scanPath string, outputDir string) {
	files, err := scanProject(scanPath)
	if err != nil {
		panic(err)
	}
	if outputDir == "" {
		outputDir = "./ModelsTS"
	}

	contextSlice, batches := parsingProject(files)

	//Building conf
	yamlStruct := YamlConf{OutputDirPath: outputDir,
		FilenameContextsMap: make(map[string][]string),
		ContextList:         contextSlice,
		BatchesInterface:    batches}

	appendEnumInYaml(files, &yamlStruct)
	for _, s := range contextSlice {
		l := make([]string, 0)
		l = append(l, strings.ToLower(s))
		yamlStruct.FilenameContextsMap[strings.Title(s)+"Model"] = l
	}

	yamlStruct.marshallAndWrite(scanPath + "/gototoConf.yaml")
}

//Update a YAML config. It keeps modifications of filenames but do not keep batches file modification.Can update directory path too.
func UpdateConfigYaml(scanPath string, outputDir string, noMoreFiles bool) {
	var currentConf YamlConf
	_, err := currentConf.GetYamlConfig(scanPath)
	Check(err)

	//Modifying outputDir if option is not ""
	if outputDir != "" {
		currentConf.OutputDirPath = outputDir
	}

	//Scanning
	files, err := scanProject(scanPath)
	if err != nil {
		panic(err)
	}
	//Parsing
	contextSlice, batches := parsingProject(files)
	missing_contexts := currentConf.getMissingContexts(contextSlice)

	/*
		Handling contexts update
	*/

	if missing_contexts != nil {
		//If no more files then we add these contexts to pending and do not add filenames for new contexts
		//Note that generateTS won't work if some contexts are still pending (except if --ignore-pending)
		if noMoreFiles {
			currentConf.Pending = append(currentConf.Pending, missing_contexts...)

		} else {
			for _, mc := range missing_contexts {
				l := make([]string, 0)
				l = append(l, strings.ToLower(mc))
				currentConf.FilenameContextsMap[strings.Title(mc)+"Model"] = l
			}
		}
	}

	/*
		Handling batches update
	*/
	currentConf.BatchesInterface = batches
	currentConf.marshallAndWrite(scanPath + "/gototoConf.yaml")

}

//Parsing all files in project and return all the contexts and batches
func parsingProject(files []string) ([]string, map[string][]TsInterface) {
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
	return contextSlice, batches
}

func appendEnumInYaml(files []string, y *YamlConf) {

	for _, f := range files {
		if s := strings.SplitN(f, "\\", -1); s[len(s)-1] == "enums.go" {
			y.BatchesEnums = append(y.BatchesEnums, ParseEnumsGoFile(f))
		}
	}
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

//Get the missing contexts and append them to the conf
func (y *YamlConf) getMissingContexts(contexts []string) []string {
	var diff []string
	for _, nc := range contexts {
		found := false
		for _, cc := range y.ContextList {
			if nc == cc {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, nc)
			y.ContextList = append(y.ContextList, nc)
		}
	}
	return diff
}

func (y *YamlConf) marshallAndWrite(path string) {
	//Marshalling contexts
	yamlStructConf, err := yaml.Marshal(y)
	Check(err)
	//writing file
	fconf, err := os.Create(path)
	Check(err)
	_, err = fconf.Write(yamlStructConf)
	Check(err)
	fconf.Close()
}
