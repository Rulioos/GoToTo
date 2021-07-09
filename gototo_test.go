package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//scanner test
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
	assert.EqualValues(t, testbatches, ymlConf.BatchesInterface)

}

func TestScanProject(t *testing.T) {
	d, err := scanProject("./project_example")
	assert.EqualValues(t, err, nil)
	assert.ElementsMatch(t, d, []string{"project_example/main.go",
		"project_example/domain/user/user.go",
		"project_example/domain/product/product.go"})
}

//utils.go test
func TestSpaceStringsBuilder(t *testing.T) {
	assert.Equal(t, SpaceStringsBuilder("i am a test"), "iamatest")
	assert.Equal(t, SpaceStringsBuilder("i am a test") == "iam a t e s t", false)
}

func TestSetifyString(t *testing.T) {
	l := []string{"match", "France", "Perdu", "Suisse", "Perdu", "Perdu", "France", "football", "Suisse"}
	s := []string{"match", "France", "Perdu", "Suisse", "football"}
	SetifyString(&l)
	assert.ElementsMatch(t, l, s)
}

//test parser
func TestParseComment(t *testing.T) {
	comment_good := "@tsInterface"
	comment_good2 := "@tsInterface[context=\"Bonjour\"]"
	comment_bad := "@tsInterface[context]"

	contextG, err := ParseComment(comment_good)
	contextG2, err2 := ParseComment(comment_good2)
	contextB, err3 := ParseComment(comment_bad)

	//error asserts
	assert.Equal(t, err, nil)
	assert.Equal(t, err2, nil)
	assert.EqualError(t, err3, fmt.Sprintf("Comment %s does not match regex pattern", comment_bad))

	//context asserts
	assert.Equal(t, contextG, "")
	assert.Equal(t, contextG2, "bonjour")
	assert.Equal(t, contextB, "")

}

func Test_parseTags(t *testing.T) {

	type parseResponse struct {
		Jname     string
		Omitempty bool
	}

	buildParseResponse := func(j string) parseResponse { j, o := parseTags(j); return parseResponse{Jname: j, Omitempty: o} }

	jtag1 := "`json:\"name\"`"
	jtag2 := "`json:\",omitempty\"`"
	jtag3 := "`json:\"name,omitempty,-\"`"
	jtag4 := "`json:\"    ,omitempty\"`"

	expected_jtag1 := parseResponse{Jname: "name", Omitempty: false}
	expected_jtag2 := parseResponse{Jname: "", Omitempty: true}
	expected_jtag3 := parseResponse{Jname: "name", Omitempty: true}
	expected_jtag4 := parseResponse{Jname: "", Omitempty: false}

	assert.EqualValues(t, expected_jtag1, buildParseResponse(jtag1))
	assert.EqualValues(t, expected_jtag2, buildParseResponse(jtag2))
	assert.EqualValues(t, expected_jtag3, buildParseResponse(jtag3))
	assert.EqualValues(t, expected_jtag4, buildParseResponse(jtag4))
}
