package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

//Type that holds the name of a struct and it's fields.
type TsInterface struct {
	Name   string
	Fields []Field
}

//Type that holds the fields name,type and omitempty option in a struct.
type Field struct {
	Name      string
	Ftype     string
	Omitempty bool
}

//ParseComment parses a comment with pattern //@tsInterface[context="SomeContext"] and returns the context. If the context is not defined: returns empty string.
func ParseComment(c string) string {
	str := SpaceStringsBuilder(c)
	reg := `^@tsInterface(\[context=(\"[A-Za-z]+\")\]|)?$`
	compReg := regexp.MustCompile(reg)
	isNotMatch := !compReg.MatchString(str)
	if isNotMatch {
		fmt.Println("error matching comment: ", str, "\n. Does not match format: ", "@tsInterface[context=\"internal\"] or @tsInterface")
		return ""
	}
	//Checking if there is a context
	if len(str) <= 14 {
		return ""
	} else {
		//Getting context
		contextReg := "\"[A-Za-z]+\""
		compContextReg := regexp.MustCompile(contextReg)
		context := compContextReg.FindString(str)
		context = context[1 : len(context)-1]
		return strings.ToLower(context)
	}
}

//getContextsAndPos loops over the comments of parsed file and fill a set of contexts with their name and a map[line_in_file]context
func getContextsAndPos(coms []*ast.CommentGroup, fset *token.FileSet, comMap *map[int]string, contextSlice *[]string) {
	for _, s := range coms {
		context := ParseComment(s.Text())
		(*comMap)[fset.Position(s.Pos()).Line] = context
		*contextSlice = append(*contextSlice, context)
	}
}

//ParseGoFile parses a file to return batches of map [context string] []ts TsInterface for one file.
func ParseGoFile(filename string, contextSlice *[]string) map[string][]TsInterface {
	fset := token.NewFileSet()
	coms, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error reaching file:", filename, err)
		return nil
	}
	//Parsing comments
	comMap := make(map[int]string)
	batches := make(map[string][]TsInterface)
	localContextSlice := make([]string, 0)

	getContextsAndPos(coms.Comments, fset, &comMap, &localContextSlice)
	SetifyString(&localContextSlice)
	*contextSlice = append(*contextSlice, localContextSlice...)

	//Building batches
	for _, c := range localContextSlice {
		batches[c] = make([]TsInterface, 0)
	}

	for _, node := range coms.Decls {
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					line := fset.Position(typeSpec.Pos()).Line
					if c, ok := comMap[line-1]; ok {
						name := typeSpec.Name.Name
						tsinterface := TsInterface{
							Name:   name,
							Fields: make([]Field, 0),
						}
						switch typeSpec.Type.(type) {
						case *ast.StructType:
							structType := typeSpec.Type.(*ast.StructType)
							for _, field := range structType.Fields.List {
								switch field.Type.(type) {
								case *ast.Ident: // field is basetype or Object
									fieldtype := field.Type.(*ast.Ident)
									appendFields(field, &tsinterface, fieldtype.Name)
								case *ast.ArrayType: // field is []basetype or []Object
									fieldtype_name := fmt.Sprintf("%v[]", field.Type.(*ast.ArrayType).Elt)
									appendFields(field, &tsinterface, fieldtype_name)
								}

							}

						}
						batches[c] = append(batches[c], tsinterface)
					}
				}
			}
		}
	}
	return batches

}

//appendFields is used to append files to a tsInterface while parsing a file
func appendFields(field *ast.Field, tsinterface *TsInterface, fieldType string) {
	fname, omitempty := parseTags(field.Tag.Value)
	f := Field{
		Name:      fname,
		Ftype:     MapToTs(fieldType),
		Omitempty: omitempty,
	}
	tsinterface.Fields = append(tsinterface.Fields, f)
}

//parseTags parses json tags next to field to get their name and omitempty option value.
func parseTags(tag string) (string, bool) {
	var omitempty bool
	var jname string
	json_tag_reg := `^json:\"([A-za-z]+)?((,omitempty)?|)((,-)?|)\"$` //Json identifier regex
	s := SpaceStringsBuilder(tag)[1 : len(tag)-1]

	compReg := regexp.MustCompile(json_tag_reg)
	isMatch := compReg.MatchString(s)
	if !isMatch {
		return "", false
	}

	//Looking for name and omitempty
	s = strings.SplitN(s, ":", 2)[1]
	params := strings.SplitN(s[1:len(s)-1], ",", 3)
	for _, param := range params {
		if param == "omitempty" {
			omitempty = true
		} else if param != "-" && param != "omitempty" {
			jname = param
		}
	}

	return jname, omitempty
}
