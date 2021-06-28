package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateConfigYaml(t *testing.T) {
	GenerateConfigYaml("./project_example", "./output_dir")
	var ymlConf YamlConf
	ymlConf.GetYamlConfig("./project_example")

	testOutputDirPath := "./output_dir"
	testContextList := []string{"user", "product"}
	testFilenameContextsMap := map[string][]string{
		"UserModel":    {"user"},
		"ProductModel": {"product"},
	}
	testbatches := map[string][]TsInterface{
		"user": {
			{
				"User",
				[]Field{
					{Name: "person", Ftype: "Person", Omitempty: false},
					{Name: "pwd", Ftype: "string", Omitempty: false},
					{Name: "login", Ftype: "string", Omitempty: false},
				},
			},
			{
				"Person",
				[]Field{
					{Name: "name", Ftype: "string", Omitempty: false},
					{Name: "given", Ftype: "string", Omitempty: false},
					{Name: "gender", Ftype: "string", Omitempty: false},
					{Name: "phone", Ftype: "string", Omitempty: true},
					{Name: "email", Ftype: "string", Omitempty: false},
				},
			},
		},
		"product": {
			{
				"Item",
				[]Field{
					{Name: "name", Ftype: "string", Omitempty: false},
					{Name: "price", Ftype: "number", Omitempty: false},
					{Name: "in_stock", Ftype: "boolean", Omitempty: false},
				},
			},
		},
	}
	assert.Equal(t, ymlConf.OutputDirPath, testOutputDirPath)
	assert.ElementsMatch(t, ymlConf.ContextList, testContextList)
	assert.EqualValues(t, testFilenameContextsMap, ymlConf.FilenameContextsMap)
	assert.EqualValues(t, testbatches, ymlConf.Batches)

}
