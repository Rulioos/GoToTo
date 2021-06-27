package main

import (
	"fmt"
	"os"

	"github.com/Pallinder/go-randomdata"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func writeInterface(f *os.File, i []TsInterface, indent string) {
	for _, tsIf := range i {
		declaration := fmt.Sprintf(indent+"export interface %s {\n", tsIf.Name)
		_, err := f.WriteString(declaration)
		check(err)
		for _, field := range tsIf.Fields {
			omit := ""
			if field.Omitempty {
				omit = "?"
			}
			field := fmt.Sprintf(indent+"\t%s %s: %s;\n", field.Name, omit, field.Ftype)
			_, err := f.WriteString(field)
			check(err)
		}

		f.WriteString(indent + "}\n\n")

	}

}

//GENERATE ONE FILE ; CONTEXTS SEPARATED BY NAMESPACES
func GenerateAll_TS_namespace(filename string, batches map[string][]TsInterface) {
	f, err := os.Create(filename + ".ts")
	check(err)

	defer f.Close()
	for c, i := range batches {
		indent := ""
		if len(c) >= 1 {
			namespace := fmt.Sprintf("declare namespace %s {\n", c)
			_, err = f.WriteString(namespace)
			check(err)
			indent = "\t"
		}

		writeInterface(f, i, indent)

		if len(c) >= 1 {
			_, err = f.WriteString("}\n\n\n")
			check(err)
		}

	}

}

// GENERATE A FOLDER CONTAINING MULTIPLE FILES. FILESNAME ARE CONTEXTS. IF NOT DEFINE IT IS A RANDOM NAME
func GenerateAll_TS_multiple_files(output_dir string, batches map[string][]TsInterface) {
	os.RemoveAll(output_dir)
	err := os.Mkdir(output_dir, 0755)
	check(err)
	for c, i := range batches {
		filename := c
		if len(c) == 0 {
			filename = randomdata.SillyName()
		}
		filepath := fmt.Sprintf("%s/%s.ts", output_dir, filename)
		f, err := os.Create(filepath)
		check(err)
		writeInterface(f, i, "")
		f.Close()
	}

}
