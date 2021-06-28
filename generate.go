package main

import (
	"fmt"
	"os"
)

func writeInterfacesWithinNamespace(f *os.File, i []TsInterface, namespace string) {
	var indent string
	var namespaceDeclaration string
	var namespaceEnd string
	if namespace != "" {
		indent = "\t"
		namespaceDeclaration = fmt.Sprintf("declare namespace %s {\n", namespace)
		namespaceEnd = "}\n\n\n"
	}
	f.WriteString(namespaceDeclaration)
	for _, tsIf := range i {
		declaration := fmt.Sprintf(indent+"export interface %s {\n", tsIf.Name)
		_, err := f.WriteString(declaration)
		Check(err)
		for _, field := range tsIf.Fields {
			omit := ""
			if field.Omitempty {
				omit = "?"
			}
			field := fmt.Sprintf(indent+"\t%s %s: %s;\n", field.Name, omit, field.Ftype)
			_, err := f.WriteString(field)
			Check(err)
		}

		f.WriteString(indent + "}\n\n")
	}
	f.WriteString(namespaceEnd)
}

//Generate all ts files from yaml conf
func GenerateFromYaml(scanPath string) {
	var ymlConf YamlConf
	var err error
	var outputDirPath string
	var namespace string
	_, err = ymlConf.GetYamlConfig(scanPath)
	Check(err)
	if ymlConf.OutputDirPath[0:2] == "./" {
		outputDirPath = scanPath + ymlConf.OutputDirPath[1:]
	} else {
		outputDirPath = ymlConf.OutputDirPath
	}

	//Creating dir
	os.RemoveAll(outputDirPath)
	err = os.Mkdir(outputDirPath, 0755)
	Check(err)
	//looping over filenames
	for filename, contexts := range ymlConf.FilenameContextsMap {
		filepath := outputDirPath + "/" + filename + ".ts"
		f, err := os.Create(filepath)
		Check(err)

		//looping over contexts in the same filename
		for index, c := range contexts {
			if len(contexts) > 1 {
				namespace = contexts[index]
			}
			if c == "User" {
				fmt.Printf("ymlConf.Batches: %v\n", ymlConf.Batches[c])
			}
			writeInterfacesWithinNamespace(f, ymlConf.Batches[c], namespace)

		}
		f.Close()
	}

}
