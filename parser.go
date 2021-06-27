package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

type TsInterface struct {
	Name   string
	Fields map[string]string
}

func ParseComments(c string) string {
	// Checking if comment can be parsed
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
		return context
	}
}

//Parse a file to match context and structs. Return batches ( map key= context, v= interfaces)
func ParseGoFile(filename string) map[string][]TsInterface {
	fset := token.NewFileSet()
	coms, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error reaching file:", filename, err)
		return nil
	}
	comMap := make(map[int]string)
	batches := make(map[string][]TsInterface)
	contextSlice := make([]string, 0)

	for _, s := range coms.Comments {
		context := ParseComments(s.Text())
		comMap[fset.Position(s.Pos()).Line] = context
		contextSlice = append(contextSlice, context)
	}
	SetifyString(&contextSlice)

	for _, c := range contextSlice {
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
							Fields: make(map[string]string),
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

// Used to append files to ts interface while parsing
func appendFields(field *ast.Field, tsinterface *TsInterface, fieldType string) {
	fname, _ := parseTags(field.Tag.Value)
	tsinterface.Fields[fname] = MapToTs(fieldType)
}

func parseTags(tag string) (string, bool) {
	var omitempty bool
	var jname string
	json_tag_reg := `^json:\"([A-za-z]+)?((,omitempty)?|)((,-)?|)\"$`
	s := SpaceStringsBuilder(tag)[1 : len(tag)-1]

	compReg := regexp.MustCompile(json_tag_reg)
	isMatch := compReg.MatchString(s)
	if !isMatch {
		return "", false
	}

	s = strings.SplitN(s, ":", 2)[1]
	params := strings.SplitN(s[1:len(s)-1], ",", 3)
	for _, param := range params {
		if param == "omitempty" {
			omitempty = true
		} else if param != "-" && param != "omitempty" {
			jname = param
			fmt.Printf("jname: %v\n", jname)
		}
	}

	return jname, omitempty
}
